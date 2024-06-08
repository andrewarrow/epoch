# epoch

time tracker and todo list a clone of johannesjo's [super-productivity](https://github.com/johannesjo/super-productivity) in golang 

![super-productivity](https://i.imgur.com/szpWtFz.png)

# video demo
[![epoch Video Demo](https://i.imgur.com/pds2pzz.png)](https://www.youtube.com/watch?v=g2jfyMvhi34)

# screen shots

![epoch1](https://i.imgur.com/FpKfqf8.png)
![epoch2](https://i.imgur.com/hHADj3L.png)
![epoch3](https://i.imgur.com/iOvs5PS.png)
![epoch4](https://i.imgur.com/dQ8CZ0D.png)



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
