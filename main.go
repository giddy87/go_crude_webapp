package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pkg/handlers"
	"pkg/render"

	"github.com/giddy87/pkg/config"
)

const PortNumber = ":8080"

func main() {
	var app config.Appconfig
	var tc map[string]*template.Template
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.UseCache = true

	app.TemplateCache = tc
	render.NewTemplate(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		n, err := fmt.Fprintf(w, "hello world")
	//		if err != nil {
	//test
	//			fmt.Println(err)
	//		}
	//		fmt.Println(fmt.Sprintf("Number ofbytes written %d", n))
	//	})
	fmt.Println("Application starting on Port", PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
