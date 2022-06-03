package handlers

import (
	"net/http"

	"github.com/zephyrus21/gookings/pkg/config"
	"github.com/zephyrus21/gookings/pkg/models"
	"github.com/zephyrus21/gookings/pkg/renders"
)

//? repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//! creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

//! sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote-ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["name"] = "Piyush"

	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote-ip")
	stringMap["remote-ip"] = remoteIP

	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
