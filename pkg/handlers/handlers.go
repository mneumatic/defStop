package handlers

import (
	"github.com/mneumatic/defstop/models"
	config "github.com/mneumatic/defstop/pkg/configs"
	"github.com/mneumatic/defstop/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers.
var Repo *Repository

// Repository is the repository type.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Title"] = "defstop"
	render.RenderTemplate(w, r, "home.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Title"] = "About | defStop"
	render.RenderTemplate(w, r, "about.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
