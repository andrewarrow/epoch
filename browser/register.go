package browser

import (
	"fmt"
	"syscall/js"

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
	Document.Document.Call("addEventListener", "keydown", wasm.FuncOf(keyPress))
	go loadTasks()
}

func loadTasks() {
	items := wasm.DoGetItems("/task")
	tasks := Document.Id("task-list")
	tasks.Set("innerHTML", "")
	for _, item := range items {
		a := Document.RenderToNewDiv("task", item)
		fmt.Println(a)
		tasks.Call("appendChild", a)
	}
	if len(items) == 0 {
		tasks.Set("innerHTML", "There are no tasks yet.")
	}
}

func keyPress(p0 js.Value) {
	fmt.Println(p0)
	k := p0.Get("key").String()
	if k == "Meta" || k == "Shift" || k == "Control" {
		return
	}
	if k == "q" {
		go wasm.DoPost("/api/q", nil)
	} else if k == "r" {
		Global.Location.Reload()
	}
}
