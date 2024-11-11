package obsidian

import (
	"encoding/hex"
	"hash/md5"
	"io"
	"os"
	"strings"

	. "github.com/savabush/obsidian-sync/internal/config"
)

func RemoveUselessDirs() {
	Logger.Info("Remove useless dirs")
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
	Logger.Info("Remove useless dirs done")

}

func RemoveObsidianDirIfExists() {
	Logger.Info("Check existing obsidian folder")
	if _, err := os.Stat("obsidian"); err == nil {
		if os.IsNotExist(err) {
			if err != nil {
				Logger.Fatal(err)
			}
		} else {
			Logger.Info("Remove existing obsidian folder")
			err := os.RemoveAll("obsidian")
			if err != nil {
				Logger.Fatal(err)
			}

		}
	}

}

func GetFileMD5(filepath string) (string, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	md5Checksum := hex.EncodeToString(h.Sum(nil))
	Logger.Infof("MD5: ", md5Checksum)
	return md5Checksum, nil
}
