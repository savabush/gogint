package minio

import (
	"context"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	. "github.com/savabush/obsidian-sync/internal/config"
)

type MinIO struct {
	client *minio.Client
	ctx    context.Context
}

func NewMinio() *MinIO {
	Logger.Println("Init minio client")
	Logger.Printf("Endpoint: %v", Settings.Minio.ENDPOINT)
	client, err := minio.New(
		Settings.Minio.ENDPOINT,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				Settings.Minio.ACCESS_KEY,
				Settings.Minio.SECRET_KEY,
				"",
			),
			Secure: false,
		})
	if err != nil {
		Logger.Fatal(err)
	}
	ctx := context.Background()
	Logger.Println("Init minio client done")
	return &MinIO{
		client: client,
		ctx:    ctx,
	}
}

func (m *MinIO) CreateDirIfNotExists(bucket string, filepath string) {
	Logger.Printf("Create dir %v", filepath)

	// TODO: Creating dir for files if not exists in minio

}
func (m *MinIO) UploadFiles(bucket string, pathToFiles string) {
	Logger.Printf("Upload files from dir %v", pathToFiles)

	entries, err := os.ReadDir(pathToFiles)
	if err != nil {
		Logger.Fatal(err)
	}
	for _, postsDir := range entries {
		if postsDir.IsDir() {
			entries, err := os.ReadDir(pathToFiles + "/" + postsDir.Name())
			if err != nil {
				Logger.Fatal(err)
			}
			m.CreateDirIfNotExists(bucket, postsDir.Name())
			for _, entryFile := range entries {

				if entryFile.IsDir() && entryFile.Name() == "Resources" {
					// TODO: Add handling of resources of post/article
					// idea is saving it by name "resources-{nameOfPost}-{nameOfImage}.png to bucket"
					continue
				}
				// TODO: Add checking MD5 (maybe) before upload
				filepath := pathToFiles + "/" + postsDir.Name() + "/" + entryFile.Name()

				opts := minio.PutObjectOptions{
					UserMetadata: map[string]string{
						"isPosted":     "false",
						"isTranslated": "false",
						"savedOnCloud": "false",
						"isSummarized": "false",
						// "MD5":          "",
					},
					ContentLanguage: "ru-RU",
					ContentType:     "text/markdown",
				}

				m.UploadFile(bucket, filepath, opts)
			}
		}
	}

}

func (m *MinIO) UploadFile(bucket string, filepath string, opts minio.PutObjectOptions) {
	Logger.Printf("Upload file %v", filepath)

	paths := strings.Split(filepath, "/")
	filename := paths[len(paths)-1]
	info, err := m.client.FPutObject(m.ctx, bucket, filename, filepath, opts)
	if err != nil {
		Logger.Fatal(err)
	}
	Logger.Printf("File uploaded %v", info)

}

func (m *MinIO) GetFileInfo(bucket string, filename string) {
	Logger.Printf("Get file info %v", filename)
	info, err := m.client.StatObject(m.ctx, bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		Logger.Error(err)
		return
	}
	Logger.Printf("File info %v", info)
}

func (m *MinIO) CheckFileExist(filename string) bool {
	Logger.Printf("Check file %v exist", filename)

	// exists, err := m.client.Objec(filename)
	// if err != nil {
	// 	Logger.Fatal(err)
	// }
	// return exists
	return true
}
