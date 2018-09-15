// 日志模块测试模块
package alog_test

import (
	"github.com/eastsirius/aslgo/internal/test"
	"github.com/eastsirius/aslgo/alog"
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

	count := 0
	for {
		count++
		alog.ErrorPrint("log test", count)
		time.Sleep(time.Second)
	}
}
