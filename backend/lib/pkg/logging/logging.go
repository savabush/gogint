package logging

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type MyFormatter struct{}

var levelList = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

// Format implements the logrus.Formatter interface.
// It formats log entries with the following pattern:
// "timestamp - filename - [line:number] - LEVEL - message"
//
// Parameters:
//   - entry: A pointer to the logrus.Entry to be formatted
//
// Returns:
//   - []byte: The formatted log entry as a byte slice
//   - error: Any error encountered during formatting (always nil in this implementation)
func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	level := levelList[int(entry.Level)]
	strList := strings.Split(entry.Caller.File, "/")
	fileName := strList[len(strList)-1]
	b.WriteString(fmt.Sprintf("%s - %s - [line:%d] - %s - %s\n",
		entry.Time.Format("2006-01-02 15:04:05,678"), fileName,
		entry.Caller.Line, level, entry.Message))
	return b.Bytes(), nil
}

// MakeLogger creates and configures a new logrus.Logger instance.
//
// Parameters:
//   - filename: The path to the log file where logs will be written.
//   - display: A boolean flag indicating whether to display logs on stdout in addition to writing to file.
//
// Returns:
//   - *logrus.Logger: A configured logger instance.
//
// The function creates a log file with the given filename (or opens it if it already exists),
// sets up the logger to write to this file, and optionally to stdout as well.
// It also configures the logger to use a custom formatter (MyFormatter) and to report the caller.
func MakeLogger(filename string, display bool) *logrus.Logger {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}
	logger := logrus.New()
	if display {
		logger.SetOutput(io.MultiWriter(os.Stdout, f))
	} else {
		logger.SetOutput(io.MultiWriter(f))
	}
	logger.SetReportCaller(true)
	logger.SetFormatter(&MyFormatter{})
	return logger
}
