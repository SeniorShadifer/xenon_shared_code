package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(path string) (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{Filename: path,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7})
	consoleWriteSyncer := os.Stdout

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileWriteSyncer, zap.InfoLevel),
		zapcore.NewCore(encoder, consoleWriteSyncer, zap.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return logger, nil
}
