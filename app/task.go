package app

import (
	"time"

	"github.com/andrewarrow/feedback/router"
)

func Task(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleTasks(c, false)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleProjectTasks(c, second, false)
		return
	}
	if second == "completed" && third == "" && c.Method == "GET" {
		handleTasks(c, true)
		return
	}
	if second != "" && third == "completed" && c.Method == "GET" {
		handleProjectTasks(c, second, true)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleTaskCreate(c)
		return
	}
	if second != "" && third == "" && c.Method == "POST" {
		handleTaskCreateProject(c, second)
		return
	}
	if second != "" && third == "complete" && c.Method == "POST" {
		handleTaskComplete(c, second)
		return
	}
	c.NotFound = true
}

func handleTasks(c *router.Context, completed bool) {
	send := map[string]any{}
	sql := "where scheduled_at is not null"
	if completed {
		sql = "where scheduled_at is not null and completed_at is not null"
	}
	items := c.All("task", sql+" order by created_at desc", "")
	c.DecorateList(items)
	send["items"] = items
	c.SendContentAsJson(send, 200)
}

func getProjectWithDefault(c *router.Context, guid string) map[string]any {
	p := c.One("project", "where guid=$1", guid)
	if guid == "inbox" {
		p = map[string]any{"id": 0}
	}
	return p
}
func handleProjectTasks(c *router.Context, guid string, completed bool) {
	send := map[string]any{}
	p := getProjectWithDefault(c, guid)
	sql := "where project_id=$1"
	if completed {
		sql = "where project_id=$1 and completed_at is not null"
	}
	items := c.All("task", sql+" order by created_at desc", "", p["id"])
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
func handleTaskCreateProject(c *router.Context, guid string) {
	c.ReadJsonBodyIntoParams()
	send := map[string]any{}
	c.Params["name"] = c.Params["new-task"]
	p := getProjectWithDefault(c, guid)
	c.Params["project_id"] = p["id"]
	c.ValidateCreate("task")
	c.Insert("task")
	c.SendContentAsJson(send, 200)
}

func handleTaskComplete(c *router.Context, guid string) {
	send := map[string]any{}
	c.FreeFormUpdate("update tasks set completed_at=$1 where guid=$2", time.Now(), guid)
	c.SendContentAsJson(send, 200)
}
