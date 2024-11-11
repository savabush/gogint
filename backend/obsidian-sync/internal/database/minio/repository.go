package minio

import (
	"context"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	. "github.com/savabush/obsidian-sync/internal/config"
	. "github.com/savabush/obsidian-sync/internal/services"
)

type MinIO struct {
	client *minio.Client
	ctx    context.Context

	defaultPutOpts minio.PutObjectOptions
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
		client:         client,
		ctx:            ctx,
		defaultPutOpts: opts,
	}
}

func (m *MinIO) handleFiles(bucket string, fileLocal string, recursiveCount uint8) {
	Logger.Infof("Handle dir %v", fileLocal)
	entries, err := os.ReadDir(fileLocal)
	if err != nil {
		Logger.Fatal(err)
	}
	paths := strings.Split(fileLocal, "/")
	fileInMinio := strings.Join(paths[len(paths)-int(recursiveCount)-1:], "/")
	for _, entry := range entries {
		if entry.IsDir() {
			// Saving resources by name "{nameOfPost}/Resources/{nameOfImage}.png to bucket"
			resourceFilepath := fileLocal + "/" + entry.Name() // entry.Name() must be equal to "Resources"
			recursiveCount += 1
			m.handleFiles(bucket, resourceFilepath, recursiveCount)
			continue
		}
		fileInMinio = fileInMinio + "/" + entry.Name()
		fileLocal := fileLocal + "/" + entry.Name()

		fileExistsInMinio, err := m.CheckFileExist(bucket, fileLocal, fileInMinio)
		if err != nil {
			Logger.Warn(err)
		}
		if fileExistsInMinio {
			continue
		}

		m.UploadFile(bucket, fileLocal, fileInMinio)
	}
}

func (m *MinIO) UploadFiles(bucket string, pathToFiles string) {
	Logger.Infof("Upload files from dir %v", pathToFiles)
	entries, err := os.ReadDir(pathToFiles)
	if err != nil {
		Logger.Fatal(err)
	}
	for _, postsDir := range entries {
		if postsDir.IsDir() {
			m.handleFiles(bucket, pathToFiles+"/"+postsDir.Name(), 0)
		}
	}
}

func (m *MinIO) UploadFile(bucket string, filepath string, filename string) {
	Logger.Infof("Upload file %v", filepath)

	// TODO: add MD5 hash to metadata "Checksum-MD5" key

	info, err := m.client.FPutObject(m.ctx, bucket, filename, filepath, m.defaultPutOpts)
	if err != nil {
		Logger.Fatal(err)
	}
	Logger.Infof("File uploaded %v", info.ChecksumCRC32C)

}

func (m *MinIO) GetFileChecksum(bucket string, filename string) (string, error) {
	Logger.Infof("Get checksum for file %v", filename)
	info, err := m.client.GetObject(m.ctx, bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}
	defer info.Close()
	infoStat, err := info.Stat()
	if err != nil {
		return "", err
	}

	Logger.Infof("Stat %#v", infoStat)
	return infoStat.UserMetadata["Checksum-MD5"], nil
}

func (m *MinIO) CheckFileExist(bucket string, filepath string, filename string) (bool, error) {

	Logger.Infof("Check existing file %v in minio and check checksum", filename)

	if checksum, err := m.GetFileChecksum(bucket, filename); err == nil {
		md5File, err := GetFileMD5(filepath)
		Logger.Info(checksum)
		Logger.Info(md5File)
		if err != nil {
			return false, err
		}
		if checksum == md5File {
			Logger.Infof("File %v already uploaded", filename)
			return true, nil
		}
	} else {
		return false, err
	}

	return false, nil
}
