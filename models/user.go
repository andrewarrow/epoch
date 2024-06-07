package models

import (
	"fmt"

	m "github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/router"
)

type User struct {
	m.BaseModel
	c *router.Context
}

func NewUser(c *router.Context) *User {
	u := User{}
	u.Item = c.User
	u.c = c
	return &u
}

func (u *User) Set(name, value string) string {
	u.c.Params = map[string]any{}
	u.c.Params[name] = value
	msg := u.c.Update("user", "where id=", u.Item["id"])
	fmt.Println("msg", msg)
	return msg
}
