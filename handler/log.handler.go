package handler

import (
	"github.com/mrtrom/go-graphql-example-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger func
func NewLogger(config *config.Config) *zap.SugaredLogger {
	var logLevel zapcore.Level
	switch config.Log.Level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	}

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      config.Log.OutputPaths,
		ErrorOutputPaths: config.Log.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	return sugar
}
