package logger

import (
	"log"
	"os"
)

// Logger wraps the standard logger
type Logger struct {
	*log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	l.Printf("[INFO] %s", msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.Printf("[ERROR] %s", msg)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string) {
	l.Printf("[WARN] %s", msg)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string) {
	l.Printf("[DEBUG] %s", msg)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string) {
	l.Fatalf("[FATAL] %s", msg)
}

