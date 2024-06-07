package app

import (
	"os"

	"github.com/andrewarrow/feedback/router"
)

func HandleApi(c *router.Context, second, third string) {
	if second == "q" && third == "" && c.Method == "POST" {
		os.Exit(0)
		return
	}
	c.NotFound = true
}

func handleConnect(c *router.Context) {
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
