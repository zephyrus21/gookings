package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/zephyrus21/gookings/pkg/config"
	"github.com/zephyrus21/gookings/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

//? sets the config for the template cache
func NewTemplates(a *config.AppConfig) {
	app = a
}

//! renders a template to the response writer with the given template name
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		//# gets the template cache from the app config
		tc = app.TemplateCache
	} else {
		//# rebuilds the cache
		tc, _ = CreateTemplateCache()
	}

	//# gets the template to be rendered from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalf("The template %s does not exist\n", tmpl)
	}

	//# creates a buffer to write the template to
	buf := new(bytes.Buffer)

	//# renders the template to the buffer
	_ = t.Execute(buf, data)

	//# writes the buffer to the response writer
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template:", err)
	}
}

//! creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) { //@ returns a map of templates and an error
	myCache := map[string]*template.Template{}

	//# gets all page template files in the templates directory
	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		fmt.Println("error parsing template:")
		return myCache, err
	}

	//? loops through all the pages and parses them
	for _, page := range pages {
		name := filepath.Base(page)

		//# creates a new template from the parsed page
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("error parsing template:")
			return myCache, err
		}

		//# gets all layout template files in the templates directory
		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			fmt.Println("error parsing template:")
			return myCache, err
		}

		if len(matches) > 0 {
			//# adds the layout templates to the page template
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				fmt.Println("error parsing template:")
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
