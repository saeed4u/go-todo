package go_todo

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON),middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)
	router.Route("/api/v1/", func(r chi.Router) {
		
	})
	return router
}


