package logging

import (
	"time"

	"go.uber.org/zap"
)

func Info(msg string, desc string, start time.Time) {
	Logger.Info(msg, zap.String("description", desc), zap.String("service", "API"), zap.String("timestamp", time.Now().Format(time.RFC3339)), zap.Duration("latency", time.Now().Sub(start)))
}

func Warn(msg string, desc string, start time.Time) {
	Logger.Warn(msg, zap.String("description", desc), zap.String("service", "API"), zap.String("timestamp", time.Now().Format(time.Now().Format(time.RFC3339))), zap.Duration("latency", time.Now().Sub(start)))
}

func Error(msg string, desc string, start time.Time) {
	Logger.Error(msg, zap.String("description", desc), zap.String("service", "API"), zap.String("timestamp", time.Now().Format(time.RFC3339)), zap.Duration("latency", time.Now().Sub(start)))
}

func Fatal(msg string, desc string, start time.Time) {
	Logger.Fatal(msg, zap.String("description", desc), zap.String("service", "API"), zap.String("timestamp", time.Now().Format(time.RFC3339)), zap.Duration("latency", time.Now().Sub(start)))
}

func Panic(msg string, desc string, start time.Time) {
	Logger.Panic(msg, zap.String("description", desc), zap.String("service", "API"), zap.String("timestamp", time.Now().Format(time.RFC3339)), zap.Duration("latency", time.Now().Sub(start)))
}
