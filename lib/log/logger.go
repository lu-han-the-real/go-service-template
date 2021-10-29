package log

import (
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

const LoggerKey = "logger"

// NewStdoutLogger creates a default stdout logger
// TODO: [lhan] move to service template
func NewStdoutLogger(level logrus.Level) *logrus.Logger {
    return &logrus.Logger{
        Out:          os.Stdout,
        Hooks:        make(logrus.LevelHooks),
        Formatter:    NewTextFormatter(),
        Level:        level,
    }
}

// NewTextFormatter creates a text formatter for logger with settings
func NewTextFormatter() *logrus.TextFormatter {
    return &logrus.TextFormatter{
        FullTimestamp: true,
        TimestampFormat: time.RFC3339Nano,
    }
}