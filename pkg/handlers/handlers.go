package handlers

import (
	"fmt"
	"net/http"

	"github.com/giddy87/go_crude_webapp/pkg/config"
	"github.com/giddy87/go_crude_webapp/pkg/models"
	"github.com/giddy87/go_crude_webapp/pkg/render"
)

type Repository struct {
	App *config.Appconfig
}

var Repo *Repository

func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//	_, _ = fmt.Fprintf(w, "Welcome to Home page")
	stringmap := make(map[string]string)
	stringmap["test"] = "Hello User"
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringmap,
	})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	sum := add(2, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("Welcome to About page, User %d", sum))
}

func add(x, y int) int {
	return x + y
}
