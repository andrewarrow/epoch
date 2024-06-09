package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

func loadTasks(from string) {
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
		tasks.Set("innerHTML", "No tasks match.")
		return
	}
	afterCheck := func(val string) {
		Document.Id("t" + val).Hide()
	}
	Global.AutoClick("task", "complete", tasks, "check", afterCheck)
}

func loadCompletedTasks(from string) {
	guid := guidFromEnd()
	var items []map[string]any
	if guid != "" {
		items = wasm.DoGetItems("/task/" + guid + "/completed")
	} else {
		items = wasm.DoGetItems("/task/completed")
	}
	tasks := Document.Id("completed-task-list")
	if tasks == nil {
		return
	}
	tasks.Set("innerHTML", "")
	for _, item := range items {
		a := Document.RenderToNewDiv("task", item)
		tasks.Call("appendChild", a)
	}
	if len(items) == 0 {
		tasks.Set("innerHTML", "No completed tasks.")
		return
	}
}
