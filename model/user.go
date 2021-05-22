package model

import (
	"github.com/dgrijalva/jwt-go"
)

// UserDataJwt easyjson: json
type UserDataJwt struct {
	Sub   string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserClaims struct {
	UserDataJwt
	jwt.StandardClaims
}

// ClientJwt easyjson: json
type ClientJwt struct {
	UserId    string `json:"userId"`
	ProjectId string `json:"projectId"`
}
type ClientClaims struct {
	ClientJwt
	jwt.StandardClaims
}

// Project easyjson: json
type Project struct {
	Id   int    `json:"id" database:"id"`
	Name string `json:"name" database:"name"`
	App  string `json:"app" database:"app"`
}

// RequestBody easyjson: json
type RequestBody struct {
	WidgetsCount int     `json:"widgetsCount"`
	Project      Project `json:"project"`
}

// Response easyjson: json
type Response struct {
	WidgetsCount int       `json:"widgetsCount"`
	Projects     []Project `json:"projects"`
}
