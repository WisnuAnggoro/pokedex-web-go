package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/models"
)

var functions = template.FuncMap{}

type render struct {
	cfg *config.Config
}

type Render interface {
	RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData)
	CreateTemplateCache() (map[string]*template.Template, error)
}

func NewRender(cfg *config.Config) Render {
	return &render{
		cfg: cfg,
	}
}

// RenderTemplate renders a template
func (rd *render) RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if rd.cfg.UseTemplateCache {
		tc = rd.cfg.TemplateCache
	} else {
		tc, _ = rd.CreateTemplateCache()
	}
	fmt.Println(tc)
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func (rd *render) CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./views/templates/page.*.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./views/templates/layout.*.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./views/templates/layout.*.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
