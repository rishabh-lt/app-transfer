package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogging() (logger *zap.Logger) {
	ws := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		ws,
		zap.DebugLevel,
	)
	logger = zap.New(core, zap.AddCaller())
	return logger
}
