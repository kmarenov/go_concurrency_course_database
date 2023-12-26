package initialization

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultEncoding = "json"
const defaultLevel = zapcore.InfoLevel
const defaultOutputPath = "output.log"

func CreateLogger() (*zap.Logger, error) {
	level := defaultLevel
	output := defaultOutputPath

	loggerCfg := zap.Config{
		Encoding:    defaultEncoding,
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{output},
	}

	return loggerCfg.Build()
}
