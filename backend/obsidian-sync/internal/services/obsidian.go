package obsidian

import (
	"os"
	"strings"

	. "github.com/savabush/obsidian-sync/internal/config"
)

func RemoveUselessDirs() {
	Logger.Println("Remove useless dirs")
	entries, err := os.ReadDir("obsidian")
	if err != nil {
		Logger.Fatal(err)
	}
	countDirs := len(entries)
	for _, entry := range entries {
		if entry.IsDir() {
			if !strings.HasPrefix(entry.Name(), "06") && !strings.HasPrefix(entry.Name(), "05") {
				countDirs -= 1
				err := os.RemoveAll("obsidian/" + entry.Name())
				if err != nil {
					Logger.Fatal(err)
				}
			}
		}
	}
	if countDirs == 0 {
		panic("All dirs are removed, check git repository")
	}
	Logger.Println("Remove useless dirs done")

}

func RemoveObsidianDirIfExists() {
	Logger.Println("Check existing obsidian folder")
	if _, err := os.Stat("obsidian"); err == nil {
		if os.IsNotExist(err) {
			if err != nil {
				Logger.Fatal(err)
			}
		} else {
			Logger.Println("Remove existing obsidian folder")
			err := os.RemoveAll("obsidian")
			if err != nil {
				Logger.Fatal(err)
			}

		}
	}

}
