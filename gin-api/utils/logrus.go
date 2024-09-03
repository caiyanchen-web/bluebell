package utils

import (
	"bytes"
	"fmt"
	"gin-api/config"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type logFormatter struct {
}

// Format 实现Formatter(entry *logrus.Entry)([]byte,error)接口

func (t *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	log := config.Config.Logger
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s %d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s: %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)

	}
	return b.Bytes(), nil
}
func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                //创建一个实例
	mLog.SetOutput(os.Stdout)                           //设置输出类型
	mLog.SetReportCaller(config.Config.Logger.ShowLine) //开启返回函数名和行号
	mLog.SetFormatter(&logFormatter{})                  //设置I自定义的格式
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置最低的等级
	InitDefauleLogger()  //不主事即启用全局log
	return mLog
}

// InitDefauleLogger 全局log
func InitDefauleLogger() {
	logrus.SetOutput(os.Stdout)                           //设置输出类型
	logrus.SetReportCaller(config.Config.Logger.ShowLine) //开启返回函数名和行号
	logrus.SetFormatter(&logFormatter{})                  //设置I自定义的格式
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
