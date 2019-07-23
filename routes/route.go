package routes

import (
	"go-to/controllers"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func authRoutes() []Route {
	return []Route{{"login", "POST", controllers.createAccount}, {"register", "POST", controllers.login}}
}

func todoRoutes() []Route {
	//CRUD
	return []Route{{"/", "GET", controllers.getAllTodos}, {"/", "POST", controllers.createTodo}, {"/:todoId", "GET", controllers.getTodo}, {"update", "PATCH", controllers.updateTodo}, {"delete", "DELETE", controllers.deleteTodo}}
}
