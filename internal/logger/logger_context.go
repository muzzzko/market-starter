package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"market-starter/config"
	interror "market-starter/internal/error"
)

const (
	loggerContextKey = "loggerContextKey"
)

var (
	logger *zap.Logger
)

func NewLogger(cfg *config.Config) *zap.Logger {
	if logger == nil {
		output := []string{"stderr"}
		if cfg.LoggerPathToLogFile != "" {
			output = append(output, cfg.LoggerPathToLogFile)
		}

		zapCfg := zap.Config{
			Level: zap.NewAtomicLevelAt(zap.DebugLevel),
			Development: false,
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "TIME",
				LevelKey:       "LEVEL",
				MessageKey:     "MESSAGE",
				StacktraceKey:  "STACKTRACE",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.RFC3339TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      output,
			ErrorOutputPaths: []string{"stderr"},
			InitialFields: map[string]interface{}{
				programNameField: cfg.LoggerProgramNameField,
			},
		}

		var err error
		logger, err = zapCfg.Build()
		interror.Check(err)
	}

	return logger
}

func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, loggerContextKey, WithContext(ctx).With(fields...))
}

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}

	if logger, ok := ctx.Value(loggerContextKey).(*zap.Logger); ok {
		return logger
	}

	return logger
}
