package model

import(
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string
	Token string `json:"token"; sql:"-"`
}
