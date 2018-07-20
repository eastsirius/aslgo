// 日志主接口文件
package alog

import (
	"time"
	"strings"
	"errors"
)

const (
	LogLevel_Debug = iota;
	LogLevel_Info;
	LogLevel_Notify;
	LogLevel_Warn;
	LogLevel_Error;

	logLevel_Num;
)

var levelString = [logLevel_Num]string {
	"Debug ",
	"Info  ",
	"Notify",
	"Warn  ",
	"Error ",
}

func GetLogLevelString(level int) string {
	return levelString[level]
}

func ParseLogLevelString(level string) (int, error) {
	for i, v := range levelString {
		if strings.ToUpper(strings.TrimSpace(v)) == strings.ToUpper(level) {
			return i, nil
		}
	}

	if strings.ToUpper("warning") == strings.ToUpper(level) {
		return LogLevel_Warn, nil
	}

	return -1, errors.New("unknown log level: " + level)
}


type LogParam struct {
	Time time.Time
	Level int
	Args []interface{}
}

type LogItem struct {
	Time time.Time
	Level int
	Log string
}


var logger = NewLogger()

func LogPrint(level int, a ...interface{}) {
	logger.logPrint(level, a)
}

func DebugPrint(a  ...interface{}) {
	logger.logPrint(LogLevel_Debug, a)
}

func InfoPrint(a  ...interface{}) {
	logger.logPrint(LogLevel_Info, a)
}

func NotifyPrint(a  ...interface{}) {
	logger.logPrint(LogLevel_Notify, a)
}

func WarnPrint(a  ...interface{}) {
	logger.logPrint(LogLevel_Warn, a)
}

func ErrorPrint(a  ...interface{}) {
	logger.logPrint(LogLevel_Error, a)
}


// 停止服务
func Stop() {
	context.stop()
}

// 设置日志等级
func SetLogLevel(level int) {
	context.setLogLevel(level)
}

// 设置格式化器
func SetFormater(fmt Formater) {
	context.setFormater(fmt)
}

// 添加日志输出
func AddOutput(writer Writer) {
	context.addOutput(writer)
}

// 清空日志输出
func ClearOutpus() {
	context.clearOutpus()
}
