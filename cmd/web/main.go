package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/zephyrus21/gookings/pkg/config"
	"github.com/zephyrus21/gookings/pkg/handlers"
	"github.com/zephyrus21/gookings/pkg/models"
	"github.com/zephyrus21/gookings/pkg/renders"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started on :8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	gob.Register(models.Reservation{})

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
		return err
	}

	//# stores the template cache in the app config
	app.TemplateCache = tc
	app.UseCache = false //@ cache everytime something changes

	//# creates new repository and sets it in the app config
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//# creates the new template cache
	renders.NewTemplates(&app)

	return nil
}
