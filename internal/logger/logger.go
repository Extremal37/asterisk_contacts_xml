package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func New(level string) *Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: setLevel(level),
	}))
	log.Info(fmt.Sprintf("Logging level set to %s", level))
	return &Logger{log: log}
}

func setLevel(level string) slog.Level {
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (l *Logger) Debug(msg string) {
	l.log.Debug(msg)
}
func (l *Logger) Debugf(format string, args ...any) {
	l.log.Debug(fmt.Sprintf(format, args...))
}
func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}
func (l *Logger) Infof(format string, args ...any) {
	l.log.Info(fmt.Sprintf(format, args...))
}
func (l *Logger) Warn(msg string) {
	l.log.Warn(msg)
}
func (l *Logger) Warnf(format string, args ...any) {
	l.log.Warn(fmt.Sprintf(format, args...))
}
func (l *Logger) Error(msg string) {
	l.log.Error(msg)
}
func (l *Logger) Errorf(format string, args ...any) {
	l.log.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.log.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}
