package controllers

import (
	"github.com/saeed4u/go-todo/util"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	util.Respond(w,r,util.Message(false, "Well"))
}

var Login = func(w http.ResponseWriter, r *http.Request) {

}
var ResetPassword = func(w http.ResponseWriter, r *http.Request) {

}
