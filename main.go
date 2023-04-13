package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
	"github.com/jonasz-lasut/hackathon-may/server"
	"github.com/upper/db/v4/adapter/mysql"
)

var docs = flag.Bool("docs", false, "Generate router documentation")

var dbConnection = mysql.ConnectionURL{
	Database: "monolith",
	Host:     "localhost",
	User:     "monolith",
	Password: "arthur",
}

func main() {
	flag.Parse()
	var dbHandler server.DatabaseHandler

	// Database
	if !*docs {
		initDBConnection(&dbConnection)
		log.Println("Opening connection to mysql DB")
		sess, err := mysql.Open(dbConnection)

		if err != nil {
			log.Fatal("mysql.Open: ", err)
		}
		defer sess.Close()

		dbHandler = server.DatabaseHandler{DB: sess}
	}

	// Service
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Routes
	r.Get("/healthz", dbHandler.HealthcheckHandler)
	r.Route("/article", func(r chi.Router) {
		r.Get("/", dbHandler.ArticleListGetter)
		r.Put("/", dbHandler.ArticleCreator)
	})
	r.Route("/article/{articleID}", func(r chi.Router) {
		r.Get("/", dbHandler.ArticleGetter)
		r.Post("/", dbHandler.ArticleUpdater)
	})

	//Admin routes
	r.Mount("/admin", server.AdminRouter(dbHandler))

	// Generate documentations on -docs flag
	if *docs {
		generateDocs(r)
		return
	}

	bind := ":5000"
	log.Printf("Serving at %s", bind)
	http.ListenAndServe(bind, r)
}

/*
Misc
*/
func generateDocs(r chi.Router) {
	path := "docs/routes.md"

	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	text := docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
		ProjectPath: "github.com/jonasz-lasut/hackathon-may",
		Intro:       "Welcome to May hackathon generated docs",
	})

	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
}

/*
Init
*/
func initDBConnection(conn *mysql.ConnectionURL) {
	if db := os.Getenv("MONOLITH_DB"); db != "" {
		conn.Database = db
	}
	if host := os.Getenv("MONOLITH_DB_HOST"); host != "" {
		conn.Host = host
	}
	if usr := os.Getenv("MONOLITH_DB_USER"); usr != "" {
		conn.User = usr
	}
	if pass := os.Getenv("MONOLITH_DB_PASSWORD"); pass != "" {
		conn.Password = pass
	}
}
