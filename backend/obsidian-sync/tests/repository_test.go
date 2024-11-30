package minio

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	minio_sdk "github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/savabush/obsidian-sync/internal/database/minio"
	"github.com/savabush/obsidian-sync/internal/lib"
)

// MockMinioClient is a mock implementation of the MinIO client interface
// It provides controlled behavior for testing the Repository methods
type MockMinioClient struct {
	mock.Mock
}

// PutObject mocks the PutObject method of the MinIO client
func (m *MockMinioClient) PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64,
	opts minio_sdk.PutObjectOptions) (info minio_sdk.UploadInfo, err error) {
	args := m.Called(ctx, bucketName, objectName, reader, objectSize, opts)
	return args.Get(0).(minio_sdk.UploadInfo), args.Error(1)
}

// StatObject mocks the StatObject method of the MinIO client
func (m *MockMinioClient) StatObject(ctx context.Context, bucketName, objectName string, opts minio_sdk.StatObjectOptions) (minio_sdk.ObjectInfo, error) {
	args := m.Called(ctx, bucketName, objectName, opts)
	return args.Get(0).(minio_sdk.ObjectInfo), args.Error(1)
}

// TestUploadWithRetry tests the retry mechanism of the file upload process
// It verifies:
// - Successful upload on first attempt
// - Retry behavior on failure
// - Maximum retry attempts
// - Proper error propagation
func TestUploadWithRetry(t *testing.T) {
	cleanup := lib.SetupTestLogger(t)
	defer cleanup()

	tests := []struct {
		name          string
		uploadError   error
		maxRetries    int
		retryDelay    time.Duration
		expectedError string
	}{
		{
			name:          "successful upload on first attempt",
			uploadError:   nil,
			maxRetries:    3,
			retryDelay:    time.Millisecond,
			expectedError: "",
		},
		{
			name:          "fails all retries",
			uploadError:   errors.New("permanent error"),
			maxRetries:    3,
			retryDelay:    time.Millisecond,
			expectedError: "upload failed after 3 attempts: attempt 3 failed: failed to upload file test.txt: permanent error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(MockMinioClient)
			mockClient.On("PutObject",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Return(minio_sdk.UploadInfo{}, tt.uploadError)

			repo := &minio.Repository{
				Client:     mockClient,
				Bucket:     "test-bucket",
				MaxRetries: tt.maxRetries,
				RetryDelay: tt.retryDelay,
			}

			err := repo.UploadWithRetry(context.Background(), "test.txt", []byte("test"))

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			expectedCalls := 1
			if tt.uploadError != nil {
				expectedCalls = tt.maxRetries
			}
			mockClient.AssertNumberOfCalls(t, "PutObject", expectedCalls)
		})
	}
}

// TestUploadFiles tests the bulk file upload functionality
// It verifies:
// - Successful upload of multiple files
// - Error handling when a file upload fails
// - Proper error propagation
func TestUploadFiles(t *testing.T) {
	cleanup := lib.SetupTestLogger(t)
	defer cleanup()

	tests := []struct {
		name          string
		files         []minio.File
		uploadError   error
		expectedError string
	}{
		{
			name: "successful upload of multiple files",
			files: []minio.File{
				{Name: "file1.txt", Content: []byte("content1")},
				{Name: "file2.txt", Content: []byte("content2")},
			},
			uploadError:   nil,
			expectedError: "",
		},
		{
			name: "fails on first file",
			files: []minio.File{
				{Name: "file1.txt", Content: []byte("content1")},
			},
			uploadError:   errors.New("upload error"),
			expectedError: "failed to upload file file1.txt: failed to upload file file1.txt: upload error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(MockMinioClient)
			mockClient.On("PutObject",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Return(minio_sdk.UploadInfo{}, tt.uploadError)

			repo := &minio.Repository{
				Client: mockClient,
				Bucket: "test-bucket",
			}

			err := repo.UploadFiles(context.Background(), tt.files)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockClient.AssertNumberOfCalls(t, "PutObject", len(tt.files))
		})
	}
}

// TestCheckFileExist tests the file existence check functionality
// It verifies:
// - Successful detection of existing files
// - Proper handling of non-existent files
// - Error propagation from the MinIO client
func TestCheckFileExist(t *testing.T) {
	cleanup := lib.SetupTestLogger(t)
	defer cleanup()

	tests := []struct {
		name        string
		statError   error
		expectExist bool
	}{
		{
			name:        "file exists",
			statError:   nil,
			expectExist: true,
		},
		{
			name:        "file does not exist",
			statError:   errors.New("file not found"),
			expectExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(MockMinioClient)
			mockClient.On("StatObject",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Return(minio_sdk.ObjectInfo{}, tt.statError)

			repo := &minio.Repository{
				Client: mockClient,
				Bucket: "test-bucket",
			}

			exists, err := repo.CheckFileExist(context.Background(), "test.txt")
			if tt.statError != nil {
				assert.Error(t, err)
				assert.False(t, exists)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectExist, exists)
			}
		})
	}
}
