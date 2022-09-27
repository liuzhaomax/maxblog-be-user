package core

import (
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
)

var LoggerSet = wire.NewSet(wire.Struct(new(Logger), "*"), wire.Bind(new(ILogger), new(*Logger)))

type Logger struct{}

type ILogger interface {
	LogSuccess(funcName string)
	LogFailure(funcName string, err error)
}

func (l Logger) LogSuccess(funcName string) {
	logger.WithFields(logger.Fields{
		"成功方法": funcName,
	}).Info(FormatInfo(110))
}

func (l Logger) LogFailure(funcName string, err error) {
	logger.WithFields(logger.Fields{
		"失败方法": funcName,
	}).Info(err.Error())
}
