// Package minio provides functionality for interacting with MinIO object storage.
// It implements a repository pattern for managing file uploads and downloads,
// with support for concurrent operations, automatic retries, and proper error handling.
package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	. "github.com/savabush/obsidian-sync/internal/config"
	
)


// MinioClient defines the interface for MinIO operations
type MinioClient interface {
	PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error)
	StatObject(ctx context.Context, bucketName, objectName string, opts minio.StatObjectOptions) (minio.ObjectInfo, error)
}

// Repository handles MinIO storage operations with support for concurrent uploads,
// automatic retries, and proper error handling. It provides a high-level interface
// for interacting with MinIO storage while maintaining proper resource management
// and error handling.
type Repository struct {
	client     MinioClient
	ctx        context.Context
	bucket     string
	maxRetries int
	retryDelay time.Duration
	putOpts    minio.PutObjectOptions
	mu         sync.Mutex // Protects metadata access
}

// RepositoryConfig holds the configuration parameters for the MinIO repository.
// It includes connection details, retry settings, and content type configurations.
type RepositoryConfig struct {
	// Endpoint is the MinIO server endpoint (e.g., "localhost:9000")
	Endpoint string
	// AccessKey is the MinIO access key credential
	AccessKey string
	// SecretKey is the MinIO secret key credential
	SecretKey string
	// Bucket is the target MinIO bucket for operations
	Bucket string
	// MaxRetries is the maximum number of upload retry attempts
	MaxRetries int
	// RetryDelay is the duration to wait between retry attempts
	RetryDelay time.Duration
	// ContentLanguage specifies the content language (e.g., "ru-RU")
	ContentLanguage string
	// ContentType specifies the MIME type of the content
	ContentType string
}

// File represents a file to be uploaded to MinIO storage.
// It supports both direct content uploads and file path-based uploads.
type File struct {
	// Name is the target filename in MinIO storage
	Name string
	// Path is the local file path (optional if Content is provided)
	Path string
	// Content is the file content (optional if Path is provided)
	Content []byte
	// Metadata is optional custom metadata to attach to the file
	Metadata map[string]string
}

// NewRepository creates a new MinIO repository instance with the provided configuration.
// It initializes the MinIO client and sets up default upload options.
// Returns an error if the client initialization fails.
func NewRepository(cfg RepositoryConfig) (*Repository, error) {
	Logger.Info("Initializing MinIO repository")
	
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			cfg.AccessKey,
			cfg.SecretKey,
			"",
		),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	opts := minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"is-posted":      "false",
			"is-translated":  "false",
			"saved-on-cloud": "false",
			"is-summarized":  "false",
		},
		ContentLanguage:      cfg.ContentLanguage,
		ContentType:          cfg.ContentType,
		SendContentMd5:       true,
		DisableContentSha256: false,
	}

	repo := &Repository{
		client:     MinioClient(client),
		ctx:        context.Background(),
		bucket:     cfg.Bucket,
		maxRetries: cfg.MaxRetries,
		retryDelay: cfg.RetryDelay,
		putOpts:    opts,
	}

	Logger.Info("MinIO repository initialized successfully")
	return repo, nil
}

// UploadFile uploads a single file to MinIO storage.
// It supports both content-based and path-based uploads, with automatic MD5 checksum
// calculation and metadata handling. The function properly manages resources and
// provides detailed error information.
func (r *Repository) UploadFile(file File) error {
	r.mu.Lock()
	if file.Metadata != nil {
		r.putOpts.UserMetadata = file.Metadata
	}
	r.mu.Unlock()

	var reader io.Reader
	var size int64

	if len(file.Content) > 0 {
		reader = bytes.NewReader(file.Content)
		size = int64(len(file.Content))
	} else if file.Path != "" {
		f, err := os.Open(file.Path)
		if err != nil {
			return fmt.Errorf("failed to open file: %v", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to get file info: %v", err)
		}

		reader = f
		size = fi.Size()
	} else {
		return fmt.Errorf("either Content or Path must be provided")
	}

	Logger.Infof("Uploading file: %s", file.Name)
	info, err := r.client.PutObject(r.ctx, r.bucket, file.Name, reader, size, r.putOpts)
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}

	Logger.Infof("File uploaded successfully: %s, size: %d", file.Name, info.Size)
	return nil
}

// UploadFiles concurrently uploads multiple files from a directory to MinIO storage.
// It walks through the directory tree, uploading files while maintaining a maximum
// number of concurrent uploads. The function provides proper error aggregation and
// resource management.
func (r *Repository) UploadFiles(dirPath string) error {
	Logger.Infof("Uploading files from directory: %s", dirPath)

	var wg sync.WaitGroup
	errChan := make(chan error, 100)
	semaphore := make(chan struct{}, 5) // Limit concurrent uploads

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		wg.Add(1)
		go func(filePath, objectName string) {
			defer wg.Done()
			semaphore <- struct{}{} // Acquire
			defer func() { <-semaphore }() // Release

			file := File{
				Name: objectName,
				Path: filePath,
			}

			if err := r.uploadWithRetry(file); err != nil {
				errChan <- fmt.Errorf("failed to upload %s: %w", objectName, err)
			}
		}(path, relPath)

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	// Wait for all uploads to complete
	wg.Wait()
	close(errChan)

	// Collect any errors
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to upload some files: %v", errs)
	}

	return nil
}

// uploadWithRetry attempts to upload a file with automatic retry logic.
// It checks for file existence before each attempt and implements exponential
// backoff between retries. The function provides detailed error information
// about each retry attempt.
func (r *Repository) uploadWithRetry(file File) error {
	var lastErr error
	for attempt := 0; attempt < r.maxRetries; attempt++ {
		if attempt > 0 {
			Logger.Infof("Retry attempt %d/%d for file %s", attempt+1, r.maxRetries, file.Name)
			time.Sleep(r.retryDelay)
		}

		exists, err := r.CheckFileExists(file.Name)
		if err != nil && minio.ToErrorResponse(err).Code != "NoSuchKey" {
			lastErr = err
			continue
		}
		if exists {
			return nil
		}

		if err := r.UploadFile(file); err == nil {
			return nil
		} else {
			lastErr = err
		}
	}
	return fmt.Errorf("failed after %d attempts: %v", r.maxRetries, lastErr)
}

// CheckFileExists verifies if a file exists in the MinIO bucket.
// It returns true if the file exists, false if it doesn't exist,
// and an error if the check operation fails.
func (r *Repository) CheckFileExists(filename string) (bool, error) {
	_, err := r.client.StatObject(r.ctx, r.bucket, filename, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// SetBucket changes the target bucket for subsequent operations
func (r *Repository) SetBucket(bucket string) {
	r.bucket = bucket
}

// GetBucket returns the current bucket name (used for testing)
func (r *Repository) GetBucket() string {
	return r.bucket
}

// SetClient replaces the MinIO client (used for testing)
func (r *Repository) SetClient(client MinioClient) {
	r.client = client
}
