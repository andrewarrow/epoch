package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document
var isTaskOpen bool

func RegisterEvents() {
	NavEvents()
	if Global.Start == "welcome.html" {
		LoginEvents()
	} else if Global.Start == "create.html" {
		CreateEvents()
	} else if Global.Start == "register.html" {
	}
}

func CreateEvents() {
	a := wasm.NewAutoForm("create")
	a.Path = "/project"
	a.After = func(id int64) {
	}
}

func NavEvents() {
	Global.Event("plus", taskOpen)
	go FetchProjects()
}

func taskOpen() {
	if isTaskOpen {
		Document.Id("tasks").Hide()
		isTaskOpen = false
	} else {
		Document.Id("tasks").Show()
		Document.Id("new-task").Focus()
		isTaskOpen = true
		a := wasm.NewAutoForm("tasks")
		a.Path = "/task"
		a.After = func(id int64) {
			Document.Id("new-task").Set("value", "")
			go loadTasks()
		}
		Global.AddAutoForm(a)
	}
}

func LoginEvents() {
	go loadTasks()
}

func loadTasks() {
	items := wasm.DoGetItems("/task")
	tasks := Document.Id("task-list")
	tasks.Set("innerHTML", "")
	for _, item := range items {
		a := Document.RenderToNewDiv("task", item)
		tasks.Call("appendChild", a)
	}
	if len(items) == 0 {
		tasks.Set("innerHTML", "There are no tasks yet.")
	}
}
