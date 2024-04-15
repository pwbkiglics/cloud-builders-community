package builder

import (
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	return sugar
}

func getLogLevelFromString(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.InfoLevel
	case "warn", "warning":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

func InitLogger(logLevel string) *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(getLogLevelFromString(logLevel))
	cfg.Encoding = "console"
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("error initializing zap: %s\n", err)
	}

	sugar = logger.Sugar()
	sugar.Infof("Log level is set to: %s (%d)", logLevel, cfg.Level.Level())
	return sugar
}
