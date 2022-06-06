package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/zephyrus21/gookings/pkg/config"
	"github.com/zephyrus21/gookings/pkg/handlers"
	"github.com/zephyrus21/gookings/pkg/renders"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	//# creates a new template cache
	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//# stores the template cache in the app config
	app.TemplateCache = tc
	app.UseCache = false //@ cache everytime something changes

	//# creates new repository and sets it in the app config
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//# creates the new template cache
	renders.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Server started on :8080")
	// http.ListenAndServe(":8080", nil)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
