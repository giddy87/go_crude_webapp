package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func OldRenderTemplate(w http.ResponseWriter, tmpl string) {
	//This requires reading two files from disk to app everytime a request is made, not efficient, rather we try to cache it
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parseing template:", err)
		return
	}

}

tc := make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error
	//Checking to see if template is present in map (cache) already
	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err !=nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	//parse the template 
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl 
	return nil

}