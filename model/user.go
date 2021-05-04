package model

import (
	"github.com/dgrijalva/jwt-go"
)

type UserDataJwt struct {
	Sub   string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserClaims struct {
	UserDataJwt
	jwt.StandardClaims
}

type Project struct {
	Id   int    `json:"id" database:"id"`
	Name string `json:"name" database:"name"`
	App  string `json:"app" database:"app"`
}

type RequestBody struct {
	WidgetsCount int     `json:"widgets_count"`
	Project      Project `json:"project"`
}

type Response struct {
	WidgetsCount int       `json:"widgets_count"`
	Projects     []Project `json:"projects" `
}
