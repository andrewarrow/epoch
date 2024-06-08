package browser

import (
	"fmt"
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document
var isMenuOpen bool

func RegisterEvents() {
	MenuEvents()
	if Global.Start == "welcome.html" {
		LoginEvents()
	} else if Global.Start == "login.html" {
	} else if Global.Start == "register.html" {
	}
}

func MenuEvents() {
	Global.Event("menu", menuOpen)
}
func menuOpen() {
	if isMenuOpen == false {
		Document.Id("left-menu").Show()
		isMenuOpen = true
	} else {
		Document.Id("left-menu").Hide()
		isMenuOpen = false
	}
}

func LoginEvents() {
	fmt.Println("h1")
	Document.Document.Call("addEventListener", "keydown", wasm.FuncOf(keyPress))
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
