package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document
var isTaskOpen bool

func RegisterEvents() {
	NavEvents()
	if Global.Start == "welcome.html" {
		LoginEvents()
	} else if Global.Start == "project_show.html" {
		ProjectShow()
	} else if Global.Start == "project_new.html" {
		ProjectNew()
	} else if Global.Start == "project_manage.html" {
		ProjectManage()
	} else if Global.Start == "register.html" {
	}
}

func ProjectNew() {
	a := wasm.NewAutoForm("create")
	a.Path = "/project"
	a.After = func(val string) {
		Global.Toast(val + ".")
		go FetchProjects()
	}
	Global.AddAutoForm(a)
}

func NavEvents() {
	Global.Event("plus", taskOpen)
	go FetchProjects()
}

func guidFromEnd() string {
	tokens := strings.Split(Global.Location.Href, "/")
	if len(tokens) > 3 {
		guid := tokens[len(tokens)-1]
		return guid
	}
	return ""
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
		guid := guidFromEnd()
		if guid != "" {
			a.Path = "/task/" + guid
		}
		a.After = func(val string) {
			Document.Id("new-task").Set("value", "")
			go loadTasks("add+")
			go FetchProjects()
		}
		Global.AddAutoForm(a)
	}
}

func LoginEvents() {
	go loadTasks("welcome")
}
