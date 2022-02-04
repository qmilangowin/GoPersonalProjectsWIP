package web

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/qmilangowin/imagebox/pkg/database/models"
)

type TemplateData struct {
	CurrentYear int
	Image       *models.Annotation
	Images      []*models.Annotation
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")

}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

//NewTemplateCache creates a new cache of templates for UI elements
func NewTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}
	//filepath.Glob will get a slice of all filepaths with the extension `page.tmpl`
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you // call the ParseFiles() method. This means we have to use template.New() to // create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// ts, err := template.ParseFiles(page)
		// if err != nil {
		// 	return nil, err
		// }

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
