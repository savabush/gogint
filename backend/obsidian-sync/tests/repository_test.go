package minio_test

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	minio_repo "github.com/savabush/obsidian-sync/internal/database/minio"
)

// MockMinioClient is a mock implementation of the MinIO client interface
type MockMinioClient struct {
	mock.Mock
}

func (m *MockMinioClient) PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64,
	opts minio.PutObjectOptions) (info minio.UploadInfo, err error) {
	args := m.Called(ctx, bucketName, objectName, reader, objectSize, opts)
	return args.Get(0).(minio.UploadInfo), args.Error(1)
}

func (m *MockMinioClient) StatObject(ctx context.Context, bucketName, objectName string, opts minio.StatObjectOptions) (minio.ObjectInfo, error) {
	args := m.Called(ctx, bucketName, objectName, opts)
	return args.Get(0).(minio.ObjectInfo), args.Error(1)
}

func setupTestRepo(t *testing.T) (*minio_repo.Repository, *MockMinioClient, func()) {
	mockClient := new(MockMinioClient)
	cfg := minio_repo.RepositoryConfig{
		Endpoint:        "test:9000",
		AccessKey:       "test",
		SecretKey:       "test",
		Bucket:         "test-bucket",
		MaxRetries:     3,
		RetryDelay:     time.Millisecond,
		ContentLanguage: "en-US",
		ContentType:    "application/octet-stream",
	}

	repo, err := minio_repo.NewRepository(cfg)
	require.NoError(t, err)

	// Replace the real client with our mock
	repo.SetClient(mockClient)

	cleanup := func() {
		mockClient.AssertExpectations(t)
	}

	return repo, mockClient, cleanup
}

func TestUploadFile(t *testing.T) {
	tests := []struct {
		name          string
		file          minio_repo.File
		setupMock     func(*MockMinioClient)
		expectedError string
	}{
		{
			name: "successful upload with content",
			file: minio_repo.File{
				Name:    "test.txt",
				Content: []byte("test content"),
			},
			setupMock: func(m *MockMinioClient) {
				m.On("PutObject",
					mock.Anything,
					"test-bucket",
					"test.txt",
					mock.Anything,
					int64(len("test content")),
					mock.Anything,
				).Return(minio.UploadInfo{}, nil)
			},
		},
		{
			name: "upload failure",
			file: minio_repo.File{
				Name:    "test.txt",
				Content: []byte("test content"),
			},
			setupMock: func(m *MockMinioClient) {
				m.On("PutObject",
					mock.Anything,
					"test-bucket",
					"test.txt",
					mock.Anything,
					int64(len("test content")),
					mock.Anything,
				).Return(minio.UploadInfo{}, errors.New("upload failed"))
			},
			expectedError: "upload failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, mockClient, cleanup := setupTestRepo(t)
			defer cleanup()

			tt.setupMock(mockClient)

			err := repo.UploadFile(tt.file)
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUploadFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "minio_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test files
	file1Path := filepath.Join(tempDir, "file1.txt")
	file2Path := filepath.Join(tempDir, "file2.txt")
	require.NoError(t, os.WriteFile(file1Path, []byte("test content 1"), 0644))
	require.NoError(t, os.WriteFile(file2Path, []byte("test content 2"), 0644))

	tests := []struct {
		name          string
		setupMock     func(*MockMinioClient)
		expectedError string
	}{
		{
			name: "successful upload of directory",
			setupMock: func(m *MockMinioClient) {
				// First file
				m.On("StatObject",
					mock.Anything,
					"test-bucket",
					"file1.txt",
					mock.Anything,
				).Return(minio.ObjectInfo{}, minio.ErrorResponse{Code: "NoSuchKey"}).Once()

				m.On("PutObject",
					mock.Anything,
					"test-bucket",
					"file1.txt",
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(minio.UploadInfo{}, nil).Once()

				// Second file
				m.On("StatObject",
					mock.Anything,
					"test-bucket",
					"file2.txt",
					mock.Anything,
				).Return(minio.ObjectInfo{}, minio.ErrorResponse{Code: "NoSuchKey"}).Once()

				m.On("PutObject",
					mock.Anything,
					"test-bucket",
					"file2.txt",
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(minio.UploadInfo{}, nil).Once()
			},
		},
		{
			name: "upload failure",
			setupMock: func(m *MockMinioClient) {
				// First file
				m.On("StatObject",
					mock.Anything,
					"test-bucket",
					"file1.txt",
					mock.Anything,
				).Return(minio.ObjectInfo{}, minio.ErrorResponse{Code: "NoSuchKey"}).Times(3)

				m.On("PutObject",
					mock.Anything,
					"test-bucket",
					"file1.txt",
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(minio.UploadInfo{}, errors.New("upload failed")).Times(3)

				// Second file
				m.On("StatObject",
					mock.Anything,
					"test-bucket",
					"file2.txt",
					mock.Anything,
				).Return(minio.ObjectInfo{}, minio.ErrorResponse{Code: "NoSuchKey"}).Times(3)

				m.On("PutObject",
					mock.Anything,
					"test-bucket",
					"file2.txt",
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(minio.UploadInfo{}, errors.New("upload failed")).Times(3)
			},
			expectedError: "failed to upload some files: [failed to upload file2.txt: failed after 3 attempts: failed to upload file: upload failed failed to upload file1.txt: failed after 3 attempts: failed to upload file: upload failed]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, mockClient, cleanup := setupTestRepo(t)
			defer cleanup()

			tt.setupMock(mockClient)

			err := repo.UploadFiles(tempDir)
			if tt.expectedError != "" {
				// Split error messages into parts and sort them for order-independent comparison
				actualParts := strings.Split(strings.TrimPrefix(err.Error(), "failed to upload some files: ["), "]")[0]
				actualErrors := strings.Split(actualParts, " ")
				sort.Strings(actualErrors)
				actualSorted := "failed to upload some files: [" + strings.Join(actualErrors, " ") + "]"

				expectedParts := strings.Split(strings.TrimPrefix(tt.expectedError, "failed to upload some files: ["), "]")[0]
				expectedErrors := strings.Split(expectedParts, " ")
				sort.Strings(expectedErrors)
				expectedSorted := "failed to upload some files: [" + strings.Join(expectedErrors, " ") + "]"

				assert.Equal(t, expectedSorted, actualSorted)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSetBucket(t *testing.T) {
	repo, _, cleanup := setupTestRepo(t)
	defer cleanup()

	// Test initial bucket
	assert.Equal(t, "test-bucket", repo.GetBucket())

	// Test changing bucket
	repo.SetBucket("new-bucket")
	assert.Equal(t, "new-bucket", repo.GetBucket())
}

func TestCheckFileExists(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*MockMinioClient)
		expectExist bool
		expectError bool
	}{
		{
			name: "file exists",
			setupMock: func(m *MockMinioClient) {
				m.On("StatObject",
					mock.Anything,
					"test-bucket",
					"test.txt",
					mock.Anything,
				).Return(minio.ObjectInfo{}, nil)
			},
			expectExist: true,
			expectError: false,
		},
		{
			name: "file does not exist",
			setupMock: func(m *MockMinioClient) {
				m.On("StatObject",
					mock.Anything,
					"test-bucket",
					"test.txt",
					mock.Anything,
				).Return(minio.ObjectInfo{}, minio.ErrorResponse{Code: "NoSuchKey"})
			},
			expectExist: false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, mockClient, cleanup := setupTestRepo(t)
			defer cleanup()

			tt.setupMock(mockClient)

			exists, err := repo.CheckFileExists("test.txt")

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectExist, exists)
		})
	}
}
