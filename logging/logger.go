package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(path string) (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileWriteSyncer, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	consoleWriteSyncer := os.Stdout

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriteSyncer, zap.InfoLevel),
		zapcore.NewCore(consoleEncoder, consoleWriteSyncer, zap.DebugLevel),
	)

	return zap.New(core, zap.AddCaller()), nil
}
