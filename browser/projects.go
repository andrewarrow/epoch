package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

func FetchProjects() {
	items := wasm.DoGetItems("/project")
	p := Document.Id("user-projects")
	p.Set("innerHTML", "")
	for _, item := range items {
		a := Document.RenderToNewDiv("project", item)
		p.Call("appendChild", a)
	}
}
