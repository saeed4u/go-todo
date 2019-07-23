package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/saeed4u/go-todo/repository"
	"github.com/saeed4u/go-todo/routes"
	"log"
	"net/http"
	"os"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)
	router.Route("/api/v1", func(r chi.Router) {
		for _, route := range routes.AuthRoutes() {
			r.Method(route.Method, route.Path, route.Handler)
		}
	})
	return router
}

func main() {
	repository.GetDB()
	err := godotenv.Load()
	if err != nil {
		log.Print("An error occurred while loading env file")
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "6000"
	}

	router := Routes()
	walkFunc := func(method string, route string,handler http.Handler, middlewares ...func(http.Handler) http.Handler)error {
		log.Printf("%s %s\n",method,route)
		return nil
	}
	_ = chi.Walk(router, walkFunc)
	log.Fatal(http.ListenAndServe(":"+port,router))
}
