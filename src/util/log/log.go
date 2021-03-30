package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"src/config"
	"time"
)

// 全局变量直接使用，logrus实现已带锁，使用样例如下：
// import  ."src/util/log"
//  Logger.Info("msg")
//  Logger.Debug("msg")
//  Logger.Warn("msg")
//  Logger.Error("msg")
//  Logger.Fatal("msg")
var Logger *logrus.Logger

func init() {
	Logger = getLogger()
	if Logger == nil {
		panic("Logger 初始化失败")
	}
	Logger.Info("Logger 初始化成功！")
}

func getLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.SetReportCaller(true)
	if config.C.Debug == true {
		logger.SetLevel(logrus.DebugLevel)
	}

	logConfig := config.C.LogConf
	//var logConfig = struct {
	//	LogPath string
	//	LogFileName string
	//}{"./log","logurs.log"}

	baseLogPath := path.Join(logConfig.LogPath, logConfig.LogFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d-%H-%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logger.Fatal(err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})

	logger.AddHook(lfHook)
	return logger
}

//Logger 自定义封装的形式
//func writeLog(fileName, funcName, errMsg, from string, err error) {
//	Logger.WithFields(logrus.Fields{
//		"package":  "package_name",
//		"file":     fileName,
//		"function": funcName,
//		"err":      err,
//		"from":     from,
//	}).Warn(errMsg)
//}
