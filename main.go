package main

import (
	"embed"
	"epoch/app"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
	webview "github.com/webview/webview_go"
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
		router.DB_FLAVOR = "sqlite"
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("epoch", embeddedFile)
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
		go r.ListenAndServe(":3000")
		webviewShow()
		return
	}

	arg := os.Args[1]

	if arg == "import" {
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "help" {
	}
}

func webviewShow() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("epoch")
	w.SetSize(969, 666, webview.HintNone)
	w.Navigate("http://localhost:3000")
	w.Run()
}
