// 测试模块路由
package test

type TestModule interface {
	DoTest()
}

var module_map = make(map[string] TestModule)

func RegisterModule(name string, module TestModule) {
	module_map[name] = module
}

func FindModule(name string) TestModule {
	return module_map[name]
}
