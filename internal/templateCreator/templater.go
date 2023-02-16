package templateCreator

import (
	"bytes"
	"fmt"
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

type TemplateCreator struct {
	Functions template.FuncMap
	appConfig *config.AppConfig
}

func New(appConfig *config.AppConfig) *TemplateCreator {
	return &TemplateCreator{
		appConfig: appConfig,
	}
}

func (t *TemplateCreator) Create(w http.ResponseWriter, r *http.Request, templateName string, td *TemplateData) {
	var tc map[string]*template.Template

	if t.appConfig.UseCache {
		tc = t.appConfig.TemplateCache
	} else {
		tc, _ = t.CreateTemplateCache()
	}

	templateFromCache, ok := tc[templateName]
	if !ok {
		log.Fatal("Couldn't get Template cache")
	}
	td.CSRFToken = nosurf.Token(r)
	buf := new(bytes.Buffer)
	_ = templateFromCache.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing templateCreator to a browser")
	}
}

// CreateTemplateCache creates a templateCreator cache as a map
func (t *TemplateCreator) CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
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

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
