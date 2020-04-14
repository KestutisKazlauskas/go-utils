package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"fmt"
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
	Printf(string, ...interface{})
}

type logger struct {
	log *zap.Logger
}

func (l *logger) Printf(msg string, v ...interface{}) {
	if len(v) == 0 {
		l.Info(msg)
	} else {
		l.Info(fmt.Sprintf(msg, v...))
	}
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