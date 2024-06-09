package app

import (
	"time"

	"github.com/andrewarrow/feedback/router"
)

func Task(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleTasks(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleProjectTasks(c, second)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleTaskCreate(c)
		return
	}
	if second != "" && third == "complete" && c.Method == "POST" {
		handleTaskComplete(c, second)
		return
	}
	c.NotFound = true
}

func handleTasks(c *router.Context) {
	send := map[string]any{}
	items := c.All("task", "where scheduled_at is not null order by created_at desc", "")
	c.DecorateList(items)
	send["items"] = items
	c.SendContentAsJson(send, 200)
}
func handleProjectTasks(c *router.Context, guid string) {
	send := map[string]any{}
	p := c.One("project", "where guid=$1", guid)
	items := c.All("task", "where project_id=$1 order by created_at desc", "", p["id"])
	c.DecorateList(items)
	send["items"] = items
	c.SendContentAsJson(send, 200)
}

func handleTaskCreate(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	send := map[string]any{}
	c.Params["name"] = c.Params["new-task"]
	c.ValidateCreate("task")
	c.Insert("task")
	c.SendContentAsJson(send, 200)
}

func handleTaskComplete(c *router.Context, guid string) {
	send := map[string]any{}
	c.FreeFormUpdate("update tasks set completed_at=$1 where guid=$2", time.Now(), guid)
	c.SendContentAsJson(send, 200)
}
