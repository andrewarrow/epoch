package browser

import (
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	if Global.Start == "welcome.html" {
		LoginEvents()
	} else if Global.Start == "login.html" {
	} else if Global.Start == "register.html" {
	}
}

func LoginEvents() {
	time.Sleep(time.Second * 3)

	go wasm.DoPost("/api/q", nil)
}
