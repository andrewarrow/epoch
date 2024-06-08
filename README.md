# epoch

time tracker and todo list


# keydown 

```
// Document.Document.Call("addEventListener", "keydown", wasm.FuncOf(keyPress))
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
```
