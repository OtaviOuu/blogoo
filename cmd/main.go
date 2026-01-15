package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	rest "github.com/otaviouu/blogoo/cmd/rest"
	"github.com/otaviouu/blogoo/internal/infra/db"
)

func main() {
	dbc := sqlx.MustConnect("sqlite3", ":memory:")

	postsRepo, err := db.NewPostRepository(dbc)

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	postsHandler := rest.NewRestPostsHandler(postsRepo)

	r.Route("/posts", func(r chi.Router) {
		r.Post("/", postsHandler.HandleListPosts)
		r.Get("/", postsHandler.HandleListPosts)
	})

	log.Println("rodando")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
