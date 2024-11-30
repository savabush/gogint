package lib

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/savabush/obsidian-sync/internal/config"
)

// TestLogger implements a simple logger for testing
type TestLogger struct {
	*log.Logger
}

func (l *TestLogger) Info(args ...interface{})                    { l.Println(args...) }
func (l *TestLogger) Infof(format string, args ...interface{})    { l.Printf(format+"\n", args...) }
func (l *TestLogger) Debug(args ...interface{})                   { l.Println(args...) }
func (l *TestLogger) Debugf(format string, args ...interface{})   { l.Printf(format+"\n", args...) }
func (l *TestLogger) Warning(args ...interface{})                 { l.Println(args...) }
func (l *TestLogger) Warningf(format string, args ...interface{}) { l.Printf(format+"\n", args...) }
func (l *TestLogger) Error(args ...interface{})                   { l.Println(args...) }
func (l *TestLogger) Errorf(format string, args ...interface{})   { l.Printf(format+"\n", args...) }
func (l *TestLogger) Fatal(args ...interface{})                   { l.Println(args...) }
func (l *TestLogger) Fatalf(format string, args ...interface{})   { l.Printf(format+"\n", args...) }
func (l *TestLogger) Warn(args ...interface{})                    { l.Println(args...) }
func (l *TestLogger) Warnf(format string, args ...interface{})    { l.Printf(format+"\n", args...) }

// Global test logger instance
var TestLog config.LoggerInterface

func init() {
	writer := io.MultiWriter(os.Stdout)
	TestLog = &TestLogger{
		Logger: log.New(writer, "TEST: ", log.Ldate|log.Ltime),
	}
}

func SetupTestLogger(t *testing.T) func() {
	// Replace the global logger
	originalLogger := config.Logger
	config.Logger = TestLog

	// Return cleanup function
	return func() {
		config.Logger = originalLogger
	}
}
