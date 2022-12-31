package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/giddy87/go_crude_webapp/path/filepath"
	"github.com/giddy87/go_crude_webapp/pkg/config"
	"github.com/giddy87/go_crude_webapp/pkg/models"
)

/*
	func OldRenderTemplate(w http.ResponseWriter, tmpl string) {
		//This requires reading two files from disk to app everytime a request is made, not efficient, rather we try to cache it
		parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parseing template:", err)
			return
		}
	}

var Old_tc = make(map[string]*template.Template)

	func SimpleRenderTemplate(w http.ResponseWriter, t string) {
		var tmpl *template.Template
		var err error
		//Checking to see if template is present in map (cache) already
		_, inMap := Old_tc[t]
		if !inMap {
			log.Println("creating template and adding to cache")
			err = SimplecreateTemplateCache(t)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println("Using cached template")
		}
		tmpl = Old_tc[t]
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	}

	func SimplecreateTemplateCache(t string) error {
		templates := []string{
			fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
		}

		//parse the template
		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			return err
		}

		Old_tc[t] = tmpl
		return nil

}
*/
var functions = template.FuncMap{}

var app *config.Appconfig

func NewTemplate(a *config.Appconfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//td = AddDefaultData(td)

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//tc, err := CreateTemplateCache()
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)
	_, _ = buf.WriteTo(w)
	//	if err != nil {
	//		fmt.Println("error writing template to browser")
	//	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
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
