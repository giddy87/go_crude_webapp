package handlers

import (
	"fmt"
	"net/http"
	"pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//	_, _ = fmt.Fprintf(w, "Welcome to Home page")
	render.RenderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	sum := add(2, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("Welcome to About page, User %d", sum))
}

func add(x, y int) int {
	return x + y
}
