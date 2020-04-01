package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var(
	Log loggerInterface
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			MessageKey: "msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	log := logger{}
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}

	Log = &log
}

type loggerInterface interface {
	Info(string)
	Error(string, error)
}

type logger struct {
	log *zap.Logger
}

func (l *logger) Info(msg string) {
	l.log.Info(msg)
	l.log.Sync()
}

func (l *logger) Error(msg string, err error) {
	if err != nil {
		l.log.Error(msg, zap.NamedError("error", err))
	} else {
		l.log.Error(msg)
	}
	l.log.Sync()
}