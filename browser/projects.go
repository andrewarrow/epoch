package browser

import (
	"fmt"

	"github.com/andrewarrow/feedback/wasm"
)

func FetchProjects() {
	items := wasm.DoGetItems("/project")
	for i, item := range items {
		a := Document.RenderToNewDiv("project", item)
		w := wasm.NewWrapper(a)
		that := w.SelectAllByClass("top")
		w = that[0]
		px := Document.Id(fmt.Sprintf("p%d", (i + 1)))
		if px == nil {
			break
		}
		px.Set("innerHTML", w.Get("innerHTML"))
		px.Show()
	}
}

func ProjectManage() {

	after := func() {
		Global.Location.Set("href", "/")
	}
	w := Document.Id("manage")
	Global.AutoDel("/project/", w, "trash", after)
}

func ProjectShow() {
	go loadTasks("show")
}
