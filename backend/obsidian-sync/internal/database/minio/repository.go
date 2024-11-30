package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	. "github.com/savabush/obsidian-sync/internal/config"
	. "github.com/savabush/obsidian-sync/internal/services"
)

// minioClientInterface defines the methods we need from minio.Client
type minioClientInterface interface {
	FPutObject(ctx context.Context, bucketName, objectName, filePath string, opts minio.PutObjectOptions) (minio.UploadInfo, error)
	GetObject(ctx context.Context, bucketName, objectName string, opts minio.GetObjectOptions) (*minio.Object, error)
}

type MinIO struct {
	Client         minioClientInterface
	Ctx            context.Context
	DefaultPutOpts minio.PutObjectOptions
}

func NewMinio() *MinIO {
	Logger.Info("Init minio client")
	Logger.Infof("Endpoint: %v", Settings.Minio.ENDPOINT)
	client, err := minio.New(
		Settings.Minio.ENDPOINT,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				Settings.Minio.ACCESS_KEY,
				Settings.Minio.SECRET_KEY,
				"",
			),
			TrailingHeaders: true,
			Secure:          false,
		})
	if err != nil {
		Logger.Fatal(err)
	}
	ctx := context.Background()

	opts := minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"is-posted":      "false",
			"is-translated":  "false",
			"saved-on-cloud": "false",
			"is-summarized":  "false",
		},
		ContentLanguage:      "ru-RU",
		ContentType:          "application/octet-stream",
		SendContentMd5:       false,
		DisableContentSha256: false,
	}

	Logger.Info("Init minio client done")
	return &MinIO{
		Client:         client,
		Ctx:            ctx,
		DefaultPutOpts: opts,
	}
}

// FileUploadJob represents a file to be uploaded
type FileUploadJob struct {
	Bucket    string
	LocalPath string
	MinioPath string
}

// UploadWithRetry attempts to upload a file with retry logic
func (m *MinIO) UploadWithRetry(job FileUploadJob, config WorkerConfig) error {
	var lastErr error
	for attempt := 0; attempt < config.MaxRetries; attempt++ {
		if attempt > 0 {
			Logger.Infof("Retry attempt %d/%d for file %s", attempt, config.MaxRetries-1, job.LocalPath)
			time.Sleep(config.RetryDelay)
		}

		fileExistsInMinio, err := m.CheckFileExist(job.Bucket, job.LocalPath, job.MinioPath)
		if err != nil {
			lastErr = err
			Logger.Warnf("Check file exist failed (attempt %d/%d): %v", attempt+1, config.MaxRetries, err)
			continue
		}
		if fileExistsInMinio {
			return nil
		}

		err = m.UploadFile(job.Bucket, job.LocalPath, job.MinioPath)
		if err == nil {
			return nil
		}
		lastErr = err
		Logger.Warnf("Upload failed (attempt %d/%d): %v", attempt+1, config.MaxRetries, err)
	}
	return fmt.Errorf("failed after %d attempts: %v", config.MaxRetries, lastErr)
}

// uploadWorker processes files from the jobs channel
func (m *MinIO) uploadWorker(jobs <-chan FileUploadJob, wg *sync.WaitGroup, config WorkerConfig) {
	defer wg.Done()
	for job := range jobs {
		if err := m.UploadWithRetry(job, config); err != nil {
			Logger.Errorf("Failed to upload %s: %v", job.LocalPath, err)
		}
	}
}

// UploadFile uploads a file to a specified MinIO bucket.
func (m *MinIO) UploadFile(bucket string, filepath string, filename string) error {
	Logger.Infof("Upload file %v", filepath)

	m.DefaultPutOpts.UserMetadata["Checksum-MD5"], _ = GetFileMD5(filepath)

	info, err := m.Client.FPutObject(m.Ctx, bucket, filename, filepath, m.DefaultPutOpts)
	if err != nil {
		return fmt.Errorf("upload failed: %w", err)
	}
	Logger.Infof("File uploaded %v", info.ChecksumCRC32C)
	return nil
}

// UploadFiles uploads all files from a specified directory to a MinIO bucket using concurrent workers
func (m *MinIO) UploadFiles(bucket string, pathToFiles string) {
	Logger.Infof("Upload files from dir %v", pathToFiles)

	config := DefaultWorkerConfig()
	Logger.Infof("Starting upload with %d workers and buffer size %d", config.NumWorkers, config.BufferSize)

	// Create a buffered channel for jobs
	jobs := make(chan FileUploadJob, config.BufferSize)

	// Start worker pool
	var wg sync.WaitGroup
	wg.Add(config.NumWorkers)
	for i := 0; i < config.NumWorkers; i++ {
		go m.uploadWorker(jobs, &wg, config)
	}

	// Walk through directory and send jobs
	err := filepath.Walk(pathToFiles, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Calculate relative path for MinIO
			relPath, err := filepath.Rel(pathToFiles, path)
			if err != nil {
				Logger.Warn(err)
				return nil
			}

			jobs <- FileUploadJob{
				Bucket:    bucket,
				LocalPath: path,
				MinioPath: relPath,
			}
		}
		return nil
	})

	if err != nil {
		Logger.Fatal(err)
	}

	// Close jobs channel and wait for workers to finish
	close(jobs)
	wg.Wait()
}

// GetFileChecksum retrieves the MD5 checksum of a file stored in a MinIO bucket.
//
// Parameters:
//   - bucket: The name of the MinIO bucket containing the file.
//   - filename: The name of the file in the bucket.
//
// Returns:
//   - string: The MD5 checksum of the file (ETag).
//   - error: An error if the operation fails, nil otherwise.
//
// This function uses the MinIO client to fetch the object's metadata and
// returns the ETag, which represents the MD5 checksum of the file.
func (m *MinIO) GetFileChecksum(bucket string, filename string) (string, error) {
	Logger.Infof("Get checksum for file %v", filename)
	info, err := m.Client.GetObject(m.Ctx, bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}
	defer info.Close()
	infoStat, err := info.Stat()
	if err != nil {
		return "", err
	}

	// ETag is checksum MD5
	return infoStat.ETag, nil
}

// CheckFileExist checks if a file exists in the MinIO bucket and compares its checksum.
//
// Parameters:
//   - bucket: The name of the MinIO bucket.
//   - filepath: The local file path.
//   - filename: The filename in the MinIO bucket.
//
// Returns:
//   - bool: True if the file exists in MinIO and has the same checksum, false otherwise.
//   - error: Any error encountered during the process.
func (m *MinIO) CheckFileExist(bucket string, filepath string, filename string) (bool, error) {
	Logger.Infof("Check existing file %v in minio and check checksum", filename)

	checksum, err := m.GetFileChecksum(bucket, filename)
	if err != nil {
		return false, err
	}

	md5File, err := GetFileMD5(filepath)
	if err != nil {
		return false, err
	}

	if checksum == md5File {
		Logger.Infof("File %v already uploaded", filename)
		return true, nil
	}

	return false, nil
}

// File represents a file to be uploaded to MinIO
type File struct {
	Name    string // Name of the file in MinIO
	Content []byte // Content of the file
}

// MinioClient defines the interface for MinIO operations
// This interface allows for easier testing through mocking
type MinioClient interface {
	PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error)
	StatObject(ctx context.Context, bucketName, objectName string, opts minio.StatObjectOptions) (minio.ObjectInfo, error)
}

// Repository handles MinIO storage operations
type Repository struct {
	Client     MinioClient   // MinIO client interface
	Bucket     string        // Target bucket for operations
	MaxRetries int           // Maximum number of upload retries
	RetryDelay time.Duration // Delay between retries
}

// UploadWithRetry attempts to upload a file with retry logic
func (r *Repository) UploadWithRetry(ctx context.Context, filename string, content []byte) error {
	var lastErr error
	for attempt := 0; attempt < r.MaxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(r.RetryDelay)
		}

		if err := r.Upload(ctx, filename, content); err == nil {
			return nil
		} else {
			lastErr = fmt.Errorf("attempt %d failed: %w", attempt+1, err)
		}
	}
	return fmt.Errorf("upload failed after %d attempts: %w", r.MaxRetries, lastErr)
}

// Upload performs a single file upload attempt
func (r *Repository) Upload(ctx context.Context, filename string, content []byte) error {
	reader := bytes.NewReader(content)
	_, err := r.Client.PutObject(
		ctx,
		r.Bucket,
		filename,
		reader,
		int64(len(content)),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("failed to upload file %s: %w", filename, err)
	}
	return nil
}

// UploadFiles uploads multiple files sequentially
func (r *Repository) UploadFiles(ctx context.Context, files []File) error {
	for _, file := range files {
		if err := r.Upload(ctx, file.Name, file.Content); err != nil {
			return fmt.Errorf("failed to upload file %s: %w", file.Name, err)
		}
	}
	return nil
}

// CheckFileExist verifies if a file exists in the MinIO bucket
func (r *Repository) CheckFileExist(ctx context.Context, filename string) (bool, error) {
	_, err := r.Client.StatObject(ctx, r.Bucket, filename, minio.StatObjectOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to check file existence %s: %w", filename, err)
	}
	return true, nil
}
