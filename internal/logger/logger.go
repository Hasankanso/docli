package logger

import (
	"log"
	"os"
)

// Logger provides consistent logging throughout the application
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "", 0),
		errorLogger: log.New(os.Stderr, "ERROR: ", 0),
		warnLogger:  log.New(os.Stdout, "WARNING: ", 0),
	}
}

// Info logs informational messages
func (l *Logger) Info(format string, args ...interface{}) {
	l.infoLogger.Printf(format, args...)
}

// Error logs error messages
func (l *Logger) Error(format string, args ...interface{}) {
	l.errorLogger.Printf(format, args...)
}

// Warning logs warning messages
func (l *Logger) Warning(format string, args ...interface{}) {
	l.warnLogger.Printf(format, args...)
}

// Success logs success messages
func (l *Logger) Success(format string, args ...interface{}) {
	l.infoLogger.Printf("SUCCESS: "+format, args...)
}

// Fatal logs error and exits the program
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.errorLogger.Printf(format, args...)
	os.Exit(1)
}

// Infof is an alias for Info for backwards compatibility
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(format, args...)
}

// Errorf is an alias for Error for backwards compatibility
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(format, args...)
}

// Global logger instance
var GlobalLogger = NewLogger()

// Package level convenience functions
func Info(format string, args ...interface{}) {
	GlobalLogger.Info(format, args...)
}

func Error(format string, args ...interface{}) {
	GlobalLogger.Error(format, args...)
}

func Warning(format string, args ...interface{}) {
	GlobalLogger.Warning(format, args...)
}

func Success(format string, args ...interface{}) {
	GlobalLogger.Success(format, args...)
}

func Fatal(format string, args ...interface{}) {
	GlobalLogger.Fatal(format, args...)
}
