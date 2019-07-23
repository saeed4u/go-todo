package routes

import (
	"github.com/saeed4u/go-todo/controllers"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func AuthRoutes() []Route {
	return []Route{{"/login", "POST", controllers.CreateAccount}, {"/register", "POST", controllers.Login}}
}

func TodoRoutes() []Route {
	return []Route{{"/", "GET", controllers.GetAllTodos}, {"/", "POST", controllers.CreateTodo}, {"/:todoId", "GET", controllers.GetTodo}, {"/update", "PATCH", controllers.UpdateTodo}, {"/delete", "DELETE", controllers.DeleteTodo}}
}
