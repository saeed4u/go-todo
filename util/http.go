package util

import (
	"github.com/go-chi/render"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	render.JSON(w, r, data)
}

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{}{"status": status, "message": message}
}
