package handlers

import (
	"fmt"
	"net/http"

	"github.com/zephyrus21/gookings/pkg/config"
	"github.com/zephyrus21/gookings/pkg/forms"
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
	renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s, end date is %s", start, end)))
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

}

func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "choose.page.tmpl", &models.TemplateData{})
}

func (m *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "book.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "room.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ShowLogin(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "room.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostShowLogin(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "room.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}
