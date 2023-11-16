package render

import (
	"bytes"
	"github.com/sigewulf/defstop/models"
	config "github.com/sigewulf/defstop/pkg/configs"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	//td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	// Declare variable of type map.
	var tc map[string]*template.Template

	// Get template cache if in production mode. Else create template cache.
	if app.Production {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache() // Ignore Error
	}

	// get template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td) //Add template data (td)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Get Pages
	pages, err := filepath.Glob("./templates/pages/*.gohtml")

	if err != nil {
		return cache, err
	}

	// range through pages
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/layouts/*.gohtml")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/layouts/*.gohtml")
			if err != nil {
				return cache, err
			}
		}

		partials, err := filepath.Glob("./templates/partials/*.gohtml")
		if err != nil {
			return cache, err
		}

		if len(partials) > 0 {
			ts, err = ts.ParseGlob("./templates/partials/*.gohtml")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
