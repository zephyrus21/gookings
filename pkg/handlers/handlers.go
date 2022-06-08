package handlers

import (
	"fmt"
	"log"
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

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}

// ReservationSummary displays the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	renders.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
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
