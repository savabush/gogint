package obsidian

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveUselessDirs(t *testing.T) {
	// Setup test directories
	err := os.MkdirAll("obsidian", 0755)
	assert.NoError(t, err)
	defer os.RemoveAll("obsidian")

	// Create test directories
	testDirs := []string{
		"obsidian/06-test",
		"obsidian/05-test",
		"obsidian/04-remove",
		"obsidian/07-remove",
	}

	for _, dir := range testDirs {
		err := os.MkdirAll(dir, 0755)
		assert.NoError(t, err)
	}

	// Run the function
	RemoveUselessDirs()

	// Check results
	entries, err := os.ReadDir("obsidian")
	assert.NoError(t, err)

	// Should only have directories starting with 06 or 05
	for _, entry := range entries {
		name := entry.Name()
		assert.True(t, entry.IsDir())
		assert.True(t, name == "06-test" || name == "05-test")
	}
}

func TestRemoveObsidianDirIfExists(t *testing.T) {
	// Setup test directory
	err := os.MkdirAll("obsidian/test", 0755)
	assert.NoError(t, err)

	// Run the function
	RemoveObsidianDirIfExists()

	// Verify directory is removed
	_, err = os.Stat("obsidian")
	assert.True(t, os.IsNotExist(err))
}

func TestGetFileMD5(t *testing.T) {
	// Create a test file with known content
	testContent := "test content for MD5"
	testFile := "test_file.txt"
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(testFile)

	// Get actual MD5
	actualMD5, err := GetFileMD5(testFile)
	assert.NoError(t, err)
	
	// Calculate MD5 of the actual content we wrote
	expectedMD5 := "19719cff6bb50a4946c9ccb00e561d51"  // MD5 of "test content for MD5"
	assert.Equal(t, expectedMD5, actualMD5)

	// Test with non-existent file
	_, err = GetFileMD5("non_existent_file.txt")
	assert.Error(t, err)
}

func TestRemoveUselessDirs_AllDirsRemoved(t *testing.T) {
	// Setup test directories
	err := os.MkdirAll("obsidian", 0755)
	assert.NoError(t, err)
	defer os.RemoveAll("obsidian")

	// Create only directories that should be removed
	testDirs := []string{
		"obsidian/01-remove",
		"obsidian/02-remove",
	}

	for _, dir := range testDirs {
		err := os.MkdirAll(dir, 0755)
		assert.NoError(t, err)
	}

	// The function should panic when all directories would be removed
	assert.Panics(t, func() {
		RemoveUselessDirs()
	})
}
