package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Project(c *router.Context, second, third string) {
	if second == "create" && third == "" && c.Method == "GET" {
		handleProjectCreate(c)
		return
	}
	c.NotFound = true
}

func handleProjectCreate(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("create.html", send, 200)
}
