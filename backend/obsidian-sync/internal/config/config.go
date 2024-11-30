package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"runtime"
)

type Config struct {
	GIT struct {
		URL       string
		CERT_PATH string
	}
	LOGGING struct {
		FILE_PATH string
	}
	APP struct {
		SCHEDULE int
	}
	Minio struct {
		ACCESS_KEY string
		SECRET_KEY string
		ENDPOINT   string
	}
}

// WorkerConfig holds the configuration for the upload worker pool
type WorkerConfig struct {
	NumWorkers      int           `yaml:"num_workers"`
	BufferSize      int           `yaml:"buffer_size"`
	MaxRetries      int           `yaml:"max_retries"`
	RetryDelay      time.Duration `yaml:"retry_delay"`
}

// DefaultWorkerConfig returns the default worker configuration
func DefaultWorkerConfig() WorkerConfig {
	return WorkerConfig{
		NumWorkers:  runtime.GOMAXPROCS(0) * 2, // 2 workers per CPU core
		BufferSize:  1000,                      // Larger buffer for more queuing
		MaxRetries:  3,                         // Number of retry attempts
		RetryDelay:  time.Second * 2,           // Delay between retries
	}
}

func InitConfig() Config {
    // Check for test environment
    if envFile := os.Getenv("ENV_FILE"); envFile != "" {
        if err := godotenv.Load(envFile); err != nil {
            panic(err)
        }
    } else {
        // Regular environment loading
        if _, err := os.Stat(".env"); err != nil {
            if os.IsNotExist(err) {
                err = godotenv.Load()
            }
        } else {
            err = godotenv.Load(".env")
        }
    }

    schedule := os.Getenv("APP_SCHEDULE")
    if schedule == "" {
        schedule = "60" // Default value
    }

    i, err := strconv.Atoi(schedule)
    if err != nil {
        panic(err)
    }

    return Config{
        GIT: struct {
            URL       string
            CERT_PATH string
        }{
            URL:       os.Getenv("GIT_URL"),
            CERT_PATH: os.Getenv("GIT_CERT_PATH"),
        },
        LOGGING: struct {
            FILE_PATH string
        }{
            FILE_PATH: os.Getenv("LOGGING_FILE_PATH"),
        },
        APP: struct {
            SCHEDULE int
        }{
            SCHEDULE: i,
        },
        Minio: struct {
            ACCESS_KEY string
            SECRET_KEY string
            ENDPOINT   string
        }{
            ACCESS_KEY: os.Getenv("MINIO_ACCESS_KEY"),
            SECRET_KEY: os.Getenv("MINIO_SECRET_KEY"),
            ENDPOINT:   os.Getenv("MINIO_ENDPOINT"),
        },
    }
}

var Settings = InitConfig()
