package main

import (
	"embed"
	"epoch/app"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%+v\n", os.Args)
	if len(os.Args) == 1 {
		router.UseLiveTemplates = false
		r := getRouter()
		go r.ListenAndServe(":3003")
		return
	}

	arg := os.Args[1]

	if arg == "import" {
	} else if arg == "seed" {
		router.DB_FLAVOR = "sqlite"
		r := router.NewRouter("epoch", embeddedFile)
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		app.Seed(r.ToContext())
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "run" {
		r := getRouter()
		r.ListenAndServe(":3000")
	}
}

func getRouter() *router.Router {
	router.BuildTag = buildTag
	router.EmbeddedTemplates = embeddedTemplates
	router.EmbeddedAssets = embeddedAssets
	r := router.NewRouter("DATABASE_URL", embeddedFile)
	r.Paths["/"] = app.HandleWelcome
	r.Paths["api"] = app.HandleApi
	//r.Paths["epoch"] = app.Epoch
	r.Paths["project"] = app.Project
	r.Paths["task"] = app.Task
	//r.Paths["login"] = app.Login
	//r.Paths["register"] = app.Register
	//r.Paths["admin"] = app.Admin
	r.Paths["markup"] = app.Markup
	r.BucketPath = "/Users/aa/bucket"
	r.NotLoggedInPath = "epoch/login"
	return r
}

/*
func webviewShow() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("epoch")
	w.SetSize(969, 666, webview.HintNone)
	w.Navigate("http://localhost:3003")
	w.Run()
}*/
