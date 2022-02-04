package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/qmilangowin/imagebox/pkg/web"
)

//ServerError is a helper method to print out server errors
func (app *ServerApplication) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.Log.PrintError(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//ClientError is a helper method to print out errors forwarded to the API by
//the client
func (app *ServerApplication) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *ServerApplication) replyHeaders(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

}

func (app *ServerApplication) addDefaultData(td *web.TemplateData, r *http.Request) *web.TemplateData {
	if td == nil {
		td = &web.TemplateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}

func (app *ServerApplication) render(w http.ResponseWriter, r *http.Request, name string, td *web.TemplateData) {
	//retrieve the appropriate template from the cache. If it doesn't exist we throw an error
	ts, ok := app.templateCache[name]
	if !ok {
		app.Log.PrintErrorf("No such template: %s, %s", w, name)
	}

	//initialize a buffer

	buf := new(bytes.Buffer)

	//execute the template and write to buffer. If an error occurs, we will know before the page is rendered.

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.Log.PrintError(w, err)
	}

	buf.WriteTo(w)
}
