package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLogger(debug bool) {
	var config zap.Config
	if debug {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Debug(name string, message string, fields ...zap.Field) {
	msg := fmt.Sprintf("[%s] %s", name, message)
	log.Debug(msg, fields...)
}

func Info(name string, message string, fields ...zap.Field) {
	msg := fmt.Sprintf("[%s] %s", name, message)
	log.Info(msg, fields...)
}

func Warn(name string, message string, fields ...zap.Field) {
	msg := fmt.Sprintf("[%s] %s", name, message)
	log.Warn(msg, fields...)
}

func Error(name string, message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		msg := fmt.Sprintf("[%s] %s", name, v.Error())
		log.Error(msg, fields...)
	case string:
		msg := fmt.Sprintf("[%s] %s", name, v)
		log.Error(msg, fields...)
	}
}

func Fatal(name string, message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		msg := fmt.Sprintf("[%s] %s", name, v.Error())
		log.Fatal(msg, fields...)
	case string:
		msg := fmt.Sprintf("[%s] %s", name, v)
		log.Fatal(msg, fields...)
	}
}
