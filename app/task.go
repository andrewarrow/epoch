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
	c.NotFound = true
}

func handleTasks(c *router.Context) {
	send := map[string]any{}
	items := c.All("task", "order by created_at desc", "")
	send["items"] = items
	c.SendContentAsJson(send, 200)
}

func handleTaskCreate(c *router.Context) {
	send := map[string]any{}
	c.Params = map[string]any{}
	c.Params["name"] = "test"
	c.ValidateCreate("task")
	c.Insert("task")
	c.SendContentAsJson(send, 200)
}
