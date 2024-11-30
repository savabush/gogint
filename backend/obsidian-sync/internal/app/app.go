// Package app implements the main application logic for Obsidian-Sync.
// It handles the synchronization between a Git repository containing Obsidian notes
// and MinIO storage, managing the entire process from cloning the repository to
// uploading files to appropriate MinIO buckets.
package app

import (
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	. "github.com/savabush/obsidian-sync/internal/config"
	. "github.com/savabush/obsidian-sync/internal/database/minio"
	. "github.com/savabush/obsidian-sync/internal/services"
)

// App is the main function of the Obsidian-Sync application.
// It performs the following steps:
//  1. Initializes a MinIO repository with proper configuration
//  2. Sets up SSH authentication for Git operations
//  3. Clones the Obsidian repository from the configured Git URL
//  4. Processes the cloned repository's directory structure
//  5. Uploads relevant files to MinIO storage
//
// The function uses environment variables for configuration (see .env file)
// and implements proper error handling and logging throughout the process.
// It also measures and logs the total execution time.
func App() {
	start := time.Now()

	// Initialize MinIO repository with configuration
	minioConfig := RepositoryConfig{
		Endpoint:        Settings.Minio.ENDPOINT,
		AccessKey:       Settings.Minio.ACCESS_KEY,
		SecretKey:       Settings.Minio.SECRET_KEY,
		MaxRetries:      3,
		RetryDelay:      time.Second * 2,
		ContentLanguage: "ru-RU",
		ContentType:     "application/octet-stream",
	}

	minioRepo, err := NewRepository(minioConfig)
	if err != nil {
		Logger.Fatalf("Failed to initialize MinIO repository: %v", err)
	}

	Logger.Infof("Starting obsidian-sync. Time start: %v", start)

	Logger.Info("Getting auth method SSH agent")
	_, err = os.Stat(Settings.GIT.CERT_PATH)
	if err != nil {
		Logger.Fatal(err)
	}

	publicKeys, err := ssh.NewPublicKeysFromFile("git", Settings.GIT.CERT_PATH, "")
	if err != nil {
		Logger.Fatal(err)
	}

	RemoveObsidianDirIfExists()

	Logger.Info("Git clone obsidian")
	_, err = git.PlainClone("obsidian", false, &git.CloneOptions{
		URL:               Settings.GIT.URL,
		Progress:          os.Stdout,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              publicKeys,
	})
	if err != nil {
		Logger.Fatal(err)
	}
	Logger.Info("Git clone obsidian done")

	RemoveUselessDirs()

	/*
		Struct of dirs in obsidian:
			05 - Posts/
				NewPost1/
					Resources/
						Image1.png
					NewPost1.md
				NewPost2/
					Resources
						Image.png
					NewPost2.md
			06 - Articles/
				NewArticle1/
					Resources/
						Image1.png
					NewArticle1.md
				NewArticle2/
					Resources
						Image.png
					NewArticle2.md

	*/
	entries, err := os.ReadDir("obsidian")
	if err != nil {
		Logger.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			switch entry.Name() {
			case Articles, Blog:
				// Extract bucket name from directory name
				bucketName := strings.ToLower(strings.Split(entry.Name(), " - ")[1])

				// Set the bucket for this upload operation
				minioRepo.SetBucket(bucketName)

				if err := minioRepo.UploadFiles("obsidian/" + entry.Name()); err != nil {
					Logger.Fatalf("Failed to upload files from %s: %v", entry.Name(), err)
				}
			}
		}
	}

	// TODO: send success status to orchestrator (GRPC)

	Logger.Infof("Done obsidian-sync. Time execution: %v", time.Since(start))
}
