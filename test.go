// 测试程序主入口
package main

import (
	"flag"
	"fmt"
	"aslgo/internal/test"
	_ "aslgo/internal/alog_test"
)

func main() {
	var module_name string
	flag.StringVar(&module_name, "m", "", "test module name")
	flag.Parse()

	module := test.FindModule(module_name)
	if module == nil {
		fmt.Print("unknown test module:", module_name)
		return
	}

	module.DoTest()
}
