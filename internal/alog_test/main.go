// 日志模块测试模块
package alog_test

import (
	"fmt"
	"aslgo/internal/test"
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
	fmt.Print("alog test")
}
