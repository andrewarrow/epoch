package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	LogoutEvents()
	if Global.Start == "welcome.html" {
		LoginEvents()
	} else if Global.Start == "login.html" {
	} else if Global.Start == "register.html" {
	}
}

func LoginEvents() {
}

func LogoutEvents() {
}
