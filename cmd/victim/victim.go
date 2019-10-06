package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/souryogurt/victim"
	"github.com/souryogurt/victim/rest"
)

var port = 5000
var databaseURL = "postgres://localhost/victim?sslmode=disable"

func init() {
	if portVar, ok := os.LookupEnv("PORT"); ok {
		if portValue, err := strconv.Atoi(portVar); err == nil {
			port = portValue
		}
	}
	flag.IntVar(&port, "port", port, "port to listen on")

	if dbURLVar, ok := os.LookupEnv("DATABASE_URL"); ok {
		databaseURL = dbURLVar
	}
	flag.StringVar(&databaseURL, "db", databaseURL, "an URL to database")
	flag.Parse()
}

func main() {
	svc := victim.NewVictim()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/tasks", func(r chi.Router) {
		r.Get("/", rest.GetAllTasks(svc))
		r.Post("/", rest.CreateTask(svc))
		r.Route("/{taskID}", func(r chi.Router) {
			r.Get("/", rest.GetTask(svc))
			r.Put("/", rest.UpdateTask(svc))
			r.Delete("/", rest.DeleteTask(svc))
		})
	})

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
