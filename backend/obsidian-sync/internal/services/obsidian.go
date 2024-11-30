package obsidian

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"

	. "github.com/savabush/obsidian-sync/internal/config"
)

// RemoveUselessDirs removes directories from the "obsidian" folder that don't start with "06" or "05".
// It logs the process, handles errors, and ensures that not all directories are removed.
//
// The function performs the following steps:
// 1. Reads the contents of the "obsidian" directory.
// 2. Iterates through each entry, removing directories that don't match the criteria.
// 3. Keeps a count of remaining directories.
// 4. Panics if all directories are removed, as a safeguard.
//
// Errors during directory reading or removal are logged as fatal.
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

// RemoveObsidianDirIfExists checks for the existence of the "obsidian" folder
// and removes it if it exists. It logs the process and handles any errors.
//
// The function performs the following steps:
// 1. Checks if the "obsidian" folder exists.
// 2. If it exists, removes the folder and its contents.
// 3. Logs fatal errors if any occur during the process.
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

// GetFileMD5 calculates the MD5 checksum of a file given its filepath.
//
// Parameters:
//   - filepath: The path to the file for which to calculate the MD5 checksum.
//
// Returns:
//   - string: The MD5 checksum of the file as a hexadecimal string.
//   - error: An error if any occurred during the process, nil otherwise.
//
// The function performs the following steps:
// 1. Opens the file specified by the filepath.
// 2. Calculates the MD5 hash of the file contents.
// 3. Converts the hash to a hexadecimal string.
// 4. Logs the calculated MD5 checksum.
// 5. Returns the MD5 checksum and any error encountered.
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
