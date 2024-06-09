package browser

import "github.com/andrewarrow/feedback/wasm"

func loadTasks() {
	items := wasm.DoGetItems("/task")
	tasks := Document.Id("task-list")
	if tasks == nil {
		return
	}
	tasks.Set("innerHTML", "")
	for _, item := range items {
		a := Document.RenderToNewDiv("task", item)
		tasks.Call("appendChild", a)
	}
	if len(items) == 0 {
		tasks.Set("innerHTML", "There are no tasks yet.")
	}
}
