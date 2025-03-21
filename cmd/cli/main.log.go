package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // Format timestamps as ISO8601
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // INFO, ERROR, etc.
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // "main.go:24"
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	// Ensure log directory exists
	if err := os.MkdirAll("./log", 0755); err != nil {
		panic(err) // Log error if directory creation fails
	}

	file, err := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // Crash if file creation fails
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stdout) // Print to console as well
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

func main() {
	// Determine log level based on environment
	logLevel := zapcore.InfoLevel
	if os.Getenv("APP_ENV") == "development" {
		logLevel = zapcore.DebugLevel
	}

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, logLevel)

	// Add caller info for better debugging
	logger := zap.New(core, zap.AddCaller())

	// Example logs
	logger.Debug("Debug log - only visible in development", zap.Int("line", 1))
	logger.Info("Info log", zap.Int("line", 2))
	logger.Warn("Warning log", zap.Int("line", 3))
	logger.Error("Error log", zap.Int("line", 4))

	// Flush log buffer before exiting
	defer logger.Sync()
}
