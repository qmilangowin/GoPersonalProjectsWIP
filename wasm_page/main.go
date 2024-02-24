package main

import (
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
	name string
}

func (c *App) Render() app.UI {
	return app.Div().Body(
		app.H1().Class("GCP Dropzone Admin Portal").Text("Build a GUI with Go 2"),
		app.P().Class("text").Text("sample"),
	)
}

func main() {
	h := App{}
	app.Route("/", &h)
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Admin",
		Description: "Admin portal",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
