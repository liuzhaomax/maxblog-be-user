package conf

import (
	"github.com/sirupsen/logrus"
	"io"
	"maxblog-be-template/internal/core"
	"os"
)

func InitLogger() {
	logFile := "golog.txt" // TODO 根据时间创建不同的日志文件，减小IO开支
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(902, err).Error())
	}
	logrus.SetLevel(logrus.InfoLevel) // Trace << Debug << Info << Warning << Error << Fatal << Panic
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: cfg.Logger.Color})
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
}
