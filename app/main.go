package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/boj/redistore"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/utahta/go-webapp-proto/app/assets"
	"github.com/utahta/go-webapp-proto/app/controller"
	"github.com/utahta/go-webapp-proto/app/lib/config"
	"github.com/utahta/go-webapp-proto/app/lib/db"
)

// develop under GOPATH

var (
	configPath = flag.String("c", "config", "config path")
	env        = flag.String("e", "dev", "environment")
)

func run() error {
	flag.Parse()

	// load config
	if err := config.Load(*env, *configPath); err != nil {
		return err
	}

	// open database
	if err := db.Open(); err != nil {
		return err
	}
	defer db.Close()

	// for session
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		return err
	}
	defer store.Close()

	// routes
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.FileServer("/public", assets.FileSystem())

	r.Route("/dummy", func(r chi.Router) {
		r.Get("/", controller.DummyIndex)
		r.Get("/search", controller.DummySearch)
	})

	return http.ListenAndServe(":8888", r)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
