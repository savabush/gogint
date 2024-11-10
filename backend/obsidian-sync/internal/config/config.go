package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

func InitConfig() Config {
	if _, err := os.Stat(".env"); err != nil {
		if os.IsNotExist(err) {
			err = godotenv.Load()
		}
	} else {
		err = godotenv.Load(".env")
	}

	i, err := strconv.Atoi(os.Getenv("APP_SCHEDULE"))
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
