package app

import (
	"github.com/andrewarrow/feedback/markup"
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
	if second == "manage" && third == "" && c.Method == "GET" {
		handleProjectManage(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleProjectShow(c, second)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleProjectCreate(c)
		return
	}
	if second != "" && third == "" && c.Method == "DELETE" {
		handleProjectDelete(c, second)
		return
	}
	c.NotFound = true
}

func handleProjectNew(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("project_new.html", send, 200)
}
func handleProjectShow(c *router.Context, guid string) {
	send := map[string]any{}
	c.SendContentInLayout("project_show.html", send, 200)
}
func handleProjectManage(c *router.Context) {
	send := map[string]any{}
	items := c.All("project", "order by created_at desc", "")
	send["items"] = items
	c.SendContentInLayout("project_manage.html", send, 200)
}

func handleProjectDelete(c *router.Context, guid string) {
	send := map[string]any{}
	c.FreeFormUpdate("delete from projects where guid=$1", guid)
	c.SendContentAsJson(send, 200)
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
	c.Params["color"] = markup.RandomColor()
	c.ValidateCreate("project")
	c.Insert("project")
	send["val"] = c.Params["name"]
	//go sound.PlaySound("done5")
	c.SendContentAsJson(send, 200)
}
