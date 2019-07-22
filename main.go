package go_todo

import(
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON),middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)
	return router
}