package main

import (
	"embed"
	"epoch/app"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
	"github.com/eiannone/keyboard"
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
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]

	if arg == "import" {
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "run" {

		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.HandleWelcome
		r.Paths["api"] = app.HandleApi
		//r.Paths["epoch"] = app.epoch
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = app.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.NotLoggedInPath = "epoch/login"
		go r.ListenAndServe(":" + os.Args[2])
		keyboard.Open()
		defer func() {
			_ = keyboard.Close()
		}()
		go func() {
			for {
				char, _, _ := keyboard.GetKey()
				if char == 'q' {
					os.Exit(0)
				}
			}
		}()
		webviewShow()
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
