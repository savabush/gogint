package config

import (
	"os"

	. "github.com/savabush/lib/pkg/logging"
)

var Logger LoggerInterface

func init() {
	logPath := Settings.LOGGING.FILE_PATH
	if logPath == "" {
		// Default to a temporary log file if no path is specified
		logPath = "/tmp/obsidian-sync.log"
	}
	if os.Getenv("ENV_FILE") != "" {
		logPath = "/tmp/obsidian-sync-test.log"
	}
	Logger = MakeLogger(logPath, true)
}

// LoggerInterface defines the interface for our logger
type LoggerInterface interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}
