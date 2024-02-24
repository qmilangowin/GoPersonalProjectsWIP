package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var web fs.FS

type resourcesFS struct {
	http.Handler
}

func newReourcesFS(fsys fs.FS) app.ResourceProvider {
	return resourcesFS{
		Handler: http.FileServer((http.FS(fsys))),
	}
}

func (resourcesFS) Package() string { return "" }
func (resourcesFS) Static() string  { return "" }
func (resourcesFS) AppWASM() string { return "/web/app.wasm" }

type App struct {
	app.Compo
}

func (a *App) OnNav(ctx app.Context) {
	fmt.Println("Navigated to page 2")
}

func (c *App) Render() app.UI {
	return app.Div().Body(
		// app.H1().Class("GCP Dropzone Admin Portal").Text("Build a GUI with Go 2"),
		// app.P().Class("text").Text("sample"),
		app.Div().Class("ui middle aligned center aligned grid").Style("height", "100vh"),
		app.Div().Style("max-width", "450px"),
		app.H2().Class("ui center aligned icon header"),
		app.I().Class("circular cloud upload icon"),
		app.H2().Class("ui teal image header"),
		app.Div().Class("content").Text("Log-in to your account"),
	)

	// return app.Div().Class("ui middle aligned center aligned grid").Style("height: 100vh"),
}

func main() {
	h := App{}
	app.Route("/", &h)
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Admin",
		Description: "Admin portal",
		Styles: []string{
			"https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.4.0/semantic.min.css",
		},
		// Scripts: []string{
		// 	"https://cdn.muicss.com/mui-0.10.3/js/mui.min.js",
		// },
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
