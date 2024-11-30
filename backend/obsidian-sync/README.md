# Obsidian Sync
[![Go Coverage](https://codecov.io/gh/savabush/obsidian-sync/branch/main/graph/badge.svg)](https://codecov.io/gh/savabush/gogint)

A service for synchronizing Obsidian notes with MinIO storage.

## Configuration

The application is configured using environment variables. Create a `.env` file in the root directory with the following settings:

```env
# Application Settings
APP_SCHEDULE=1                    # Schedule interval

# Logging Configuration
LOGGING_FILE_PATH=./obsidian-sync.log  # Path to log file (local development)

# Git Repository Settings
GIT_URL=git@github.com:savabush/obsidian.git  # Obsidian Git repository URL
GIT_CERT_PATH=./cert/id_rsa                   # SSH private key path (local development)
SSH_KNOWN_HOSTS=./cert/known_hosts            # SSH known hosts file (local development)

# MinIO Configuration
MINIO_ACCESS_KEY=your_access_key     # MinIO access key
MINIO_SECRET_KEY=your_secret_key     # MinIO secret key
MINIO_ENDPOINT=localhost:9000        # MinIO endpoint (local development)
```

For production deployment, update the paths accordingly:
```env
LOGGING_FILE_PATH=/logs/obsidian-sync.log
GIT_CERT_PATH=/cert/id_rsa
SSH_KNOWN_HOSTS=/cert/known_hosts
MINIO_ENDPOINT=minio:9000
```

## Project Structure

```
obsidian-sync/
├── internal/
│   ├── app/
│   │   └── app.go          # Main application logic
│   ├── database/
│   │   └── minio/
│   │       └── repository.go  # MinIO storage operations
│   └── config/
│       └── enums.go        # Configuration enums
├── tests/
│   └── repository_test.go  # Repository tests
└── .env                    # Environment configuration
```

## Components

### MinIO Repository

The MinIO repository (`repository.go`) provides functionality for:
- Uploading single files
- Concurrent upload of multiple files
- Automatic retry mechanism
- File existence checking
- Custom metadata support

Key features:
- Configurable retry mechanism
- Concurrent uploads with rate limiting
- Proper error handling and logging
- Support for both file content and file path uploads

### Application

The main application (`app.go`) handles:
- MinIO repository initialization
- Git repository cloning
- File synchronization between Git and MinIO
- Directory structure management

## Development

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your settings
3. Ensure you have MinIO running locally or configure remote MinIO settings
4. Run the tests: `go test ./tests/...`
5. Start the application: `go run main.go`

## Production Deployment

For production deployment:
1. Update the `.env` file with production paths and settings
2. Build the application: `go build`
3. Deploy using Docker or your preferred method

## Testing

Run tests with coverage:
```bash
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
```

View coverage in your browser:
```bash
go tool cover -html=coverage.txt
```

### Coverage Status
The project maintains a high level of test coverage to ensure reliability. Current coverage status:
- Repository package: ~83% coverage
- Focus on critical path testing
- Concurrent operation safety
- Error handling verification

The test suite includes:
- Repository unit tests
- Integration tests with MinIO
- Error handling tests
