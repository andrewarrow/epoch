package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleApi(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleConnect(c)
		return
	}
	c.NotFound = true
}

func handleConnect(c *router.Context) {
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
