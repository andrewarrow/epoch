package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Task(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleTasks(c)
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
	items := c.All("task", "order by created_at desc", "")
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
	c.ReadJsonBodyIntoParams()
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
