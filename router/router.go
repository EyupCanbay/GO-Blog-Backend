package router

import (
	"net/http"

	handlers "golang-blog-backend/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/posts", handlers.CreatePost)
	r.Get("/posts", handlers.GetPosts)

	return r
}
