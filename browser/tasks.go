package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

func loadTasks() {
	tokens := strings.Split(Global.Location.Href, "/")
	var items []map[string]any
	if len(tokens) > 3 {
		guid := tokens[len(tokens)-1]
		items = wasm.DoGetItems("/task/" + guid)
	} else {
		items = wasm.DoGetItems("/task")
	}
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
		return
	}
	afterCheck := func(val string) {
	}
	Global.AutoClick("task", "complete", tasks, "check", afterCheck)
}
