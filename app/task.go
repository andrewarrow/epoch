package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Task(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "POST" {
		handleTaskCreate(c)
		return
	}
	c.NotFound = true
}

func handleTaskCreate(c *router.Context) {
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
