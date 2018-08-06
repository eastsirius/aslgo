// 图形界面模块测试模块
package agui_test

import (
	"aslgo/internal/test"
	"net/http"
	"aslgo/agui"
)

func init() {
	test.RegisterModule("agui", NewAlogTestModule())
}


type alogTestModule struct {
}

func NewAlogTestModule() *alogTestModule {
	return &alogTestModule{}
}

func (tm *alogTestModule) DoTest() {
	http.HandleFunc("/", tm.MainPage)
	http.ListenAndServe("0.0.0.0:6487", nil)
}

func (tm *alogTestModule) MainPage(writer http.ResponseWriter, req *http.Request) {
	page := agui.NewHtml5Page("AGUI Demo")

	writer.WriteHeader(200)
	page.WritePage(writer)
}
