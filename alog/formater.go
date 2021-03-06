// 日志格式化器
package alog

import (
	"fmt"
)

// 格式化器接口
type Formater interface {
	// 格式化日志
	Format(log *LogParam) string
}


// 基本格式化器
type basicFormater struct {
}

func NewBasicFormater() Formater {
	return &basicFormater{}
}

func (bf *basicFormater) Format(log *LogParam) string {
	msg := ""
	if log.Format == "" {
		msg = fmt.Sprint(log.Args)
		msg = msg[1 : len(msg)-1]
	} else {
		msg = fmt.Sprintf(log.Format, log.Args...)
	}
	tm_string := log.Time.Format("15:04:05.000")
	return tm_string + " [" + GetLogLevelString(log.Level) + "] " + msg
}
