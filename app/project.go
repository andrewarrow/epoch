package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Project(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleProjects(c)
		return
	}
	if second == "new" && third == "" && c.Method == "GET" {
		handleProjectNew(c)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleProjectCreate(c)
		return
	}
	c.NotFound = true
}

func handleProjectNew(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("project_new.html", send, 200)
}

func handleProjects(c *router.Context) {
	send := map[string]any{}
	items := c.All("project", "order by created_at desc", "")
	send["items"] = items
	c.SendContentAsJson(send, 200)
}
func handleProjectCreate(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	send := map[string]any{}
	c.Params["name"] = c.Params["name"]
	c.ValidateCreate("project")
	c.Insert("project")
	c.SendContentAsJson(send, 200)
}
