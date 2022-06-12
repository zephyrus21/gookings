package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/zephyrus21/gookings/pkg/config"
	"github.com/zephyrus21/gookings/pkg/models"
	"github.com/zephyrus21/gookings/pkg/renders"
)

var app config.AppConfig

var session *scs.SessionManager

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	//# creates a new template cache
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//# stores the template cache in the app config
	app.TemplateCache = tc
	app.UseCache = true //@ cache everytime something changes

	//# creates new repository and sets it in the app config
	repo := NewRepo(&app)
	NewHandlers(repo)

	//# creates the new template cache
	renders.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", Repo.ChooseRoom)
	mux.Get("/book-room", Repo.BookRoom)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/user/login", Repo.ShowLogin)
	mux.Post("/user/login", Repo.PostShowLogin)
	mux.Get("/user/logout", Repo.Logout)

	fileServer := http.FileServer(http.Dir("../../static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

//! creates a new CSRF token tp all POST requests
func NoSurf(next http.Handler) http.Handler {
	//# creates a new CSRF handler
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//! loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

var functions = template.FuncMap{}
var pathToTemplates = "../../templates"

func CreateTestTemplateCache() (map[string]*template.Template, error) { //@ returns a map of templates and an error
	myCache := map[string]*template.Template{}

	//# gets all page template files in the templates directory
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		fmt.Println("error parsing template:")
		return myCache, err
	}

	//? loops through all the pages and parses them
	for _, page := range pages {
		name := filepath.Base(page)

		//# creates a new template from the parsed page
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("error parsing template:")
			return myCache, err
		}

		//# gets all layout template files in the templates directory
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.ayout.tmpl", pathToTemplates))
		if err != nil {
			fmt.Println("error parsing template:")
			return myCache, err
		}

		if len(matches) > 0 {
			//# adds the layout templates to the page template
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				fmt.Println("error parsing template:")
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
