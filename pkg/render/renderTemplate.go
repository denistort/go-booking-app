package render

import (
	"bytes"
	"fmt"
	"github.com/denistort/go-booking-app/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func Template[K comparable](w http.ResponseWriter, templateName string, td *TemplateData[K]) {
	appConfig := config.GetAppConfig()
	var tc map[string]*template.Template
	if appConfig.UseCache {
		tc = config.GetAppConfig().TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[templateName]
	if !ok {
		log.Fatal("Couldn't get Template cache")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to a browser")
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./pkg/templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
			return myCache, err
		}

		matches, err := filepath.Glob("./pkg/templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./pkg/templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
