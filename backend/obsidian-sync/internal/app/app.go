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

func App() {

	start := time.Now()

	minioClient := NewMinio()

	Logger.Printf("Starting obsidian-sync. Time start: %v", start)

	Logger.Println("Getting auth method SSH agent")
	_, err := os.Stat(Settings.GIT.CERT_PATH)
	if err != nil {
		Logger.Fatal(err)
	}

	publicKeys, err := ssh.NewPublicKeysFromFile("git", Settings.GIT.CERT_PATH, "")

	if err != nil {
		Logger.Fatal(err)
	}

	RemoveObsidianDirIfExists()

	Logger.Println("Git clone obsidian")
	_, err = git.PlainClone("obsidian", false, &git.CloneOptions{
		URL:               Settings.GIT.URL,
		Progress:          os.Stdout,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              publicKeys,
	})
	if err != nil {
		Logger.Fatal(err)
	}
	Logger.Println("Git clone obsidian done")

	RemoveUselessDirs()

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
			case Articles:
				minioClient.UploadFiles(bucketName, "obsidian/"+entry.Name())
			case Posts:
				minioClient.UploadFiles(bucketName, "obsidian/"+entry.Name())
			}
		}

		// TODO: send success status to orchestrator (GRPC)

		Logger.Printf("Done obsidian-sync. Time execution: %v", time.Since(start))
	}
}
