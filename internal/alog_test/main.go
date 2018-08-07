// 日志模块测试模块
package alog_test

import (
	"aslgo/internal/test"
	"aslgo/alog"
	"time"
)

func init() {
	test.RegisterModule("alog", NewAlogTestModule())
}


type alogTestModule struct {
}

func NewAlogTestModule() *alogTestModule {
	return &alogTestModule{}
}

func (tm *alogTestModule) DoTest() {
	alog.SetLogLevel(alog.LogLevel_Debug)
	alog.ClearOutpus()
	alog.AddOutput(alog.NewConsoleWriter())
	fout := alog.NewFileWriter()
	fout.Path = "alog_test"
	alog.AddOutput(fout)

	for {
		alog.ErrorPrint("log test")
		time.Sleep(time.Second)
	}
}
