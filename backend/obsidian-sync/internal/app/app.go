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

// App is the main function of the obsidian-sync application.
// It performs the following steps:
// 1. Initializes a MinIO client
// 2. Sets up SSH authentication for Git
// 3. Clones the Obsidian repository
// 4. Removes unnecessary directories
// 5. Uploads files to MinIO buckets based on directory structure
//
// The function uses various helper functions and services defined elsewhere
// in the application, such as NewMinio(), RemoveObsidianDirIfExists(),
// and RemoveUselessDirs().
//
// It also handles logging and error management throughout the process.
func App() {
	start := time.Now()

	// TODO: Rewrite creating instance of minio to from Repository
	minioClient := NewMinio()
	Logger.Infof("Starting obsidian-sync. Time start: %v", start)

	Logger.Info("Getting auth method SSH agent")
	_, err := os.Stat(Settings.GIT.CERT_PATH)
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
			var bucketName string
			if strings.ContainsAny(entry.Name(), " - ") {
				bucketName = strings.ToLower(strings.Split(entry.Name(), " - ")[1])
			}

			switch entry.Name() {
			case Articles, Blog:
				minioClient.UploadFiles(bucketName, "obsidian/"+entry.Name())
			}
		}
	}

	// TODO: send success status to orchestrator (GRPC)

	Logger.Infof("Done obsidian-sync. Time execution: %v", time.Since(start))
}
