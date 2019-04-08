package log

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"strings"
)

var logger log4go.Logger

const (
	InfoLog = iota
	DebugLog
	WarnLog
	ErrorLog
	CriLog
)
const (
	InitSer  = "Init Server"
	ReqParse = "Req Parsing"
)

//记录基本日志
func Info(arg0 interface{}, args ...interface{}) {
	logger.Info(arg0, args...)
}

//记录调试日志
func Debug(arg0 interface{}, args ...interface{}) {
	logger.Debug(arg0, args...)
}

//记录警告日志
func Warn(arg0 interface{}, args ...interface{}) {
	logger.Warn(arg0, args...)
}

//记录错误日志
func Error(arg0 interface{}, args ...interface{}) {
	logger.Error(arg0, args...)
}

//记录崩溃日志
func Critical(arg0 interface{}, args ...interface{}) {
	logger.Critical(arg0, args)
}

//记录日志, 会具体记录哪种action操作
func LogWithTag(logType int, actionType string, arg2 interface{}, args ...interface{}) {
	var msg string
	switch first := arg2.(type) {
	case string:
		// Use the string as a format string
		msg = fmt.Sprintf(first, args...)
	case func() string:
		// Log the closure (no other arguments used)
		msg = first()
	default:
		// Build a format string so that it will be similar to Sprint
		msg = fmt.Sprintf(fmt.Sprint(first)+strings.Repeat(" %v", len(args)), args...)
	}

	logMsg := "[" + actionType + "]: " + msg
	switch logType {
	case InfoLog:
		logger.Info(logMsg)
	case DebugLog:
		logger.Debug(logMsg)
	case WarnLog:
		logger.Warn(logMsg)
	case ErrorLog:
		logger.Error(logMsg)
	case CriLog:
		logger.Critical(logMsg)
	default:
		logger.Info(arg2, args)
	}
}

//日志框架初始化
func init() {
	logger = make(log4go.Logger)
	logger.LoadConfiguration("config/log4go.xml")
	logger.Info("[" + InitSer + "]: " + "日志框架初始化完成")
}
