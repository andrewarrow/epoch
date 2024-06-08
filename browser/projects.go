package browser

import (
	"fmt"

	"github.com/andrewarrow/feedback/wasm"
)

func FetchProjects() {
	items := wasm.DoGetItems("/project")
	for i, item := range items {
		a := Document.RenderToNewDiv("project", item)
		w := wasm.NewWrapper(a)
		px := Document.Id(fmt.Sprintf("p%d", (i + 1)))
		px.Set("innerHTML", w.Get("innerHTML"))
		px.Show()
	}
}
