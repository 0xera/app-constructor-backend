package api

import (
	"github.com/0xera/app-constructor-backend/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Service struct {
	JwtService         *auth.JwtService
	GoogleOauthService *auth.GoogleOauthService
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Hello :)")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.UserClaims)
	return c.String(http.StatusOK, "Welcome "+claims.Email+"!")
}

func (receiver *Service) Serve() {

	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", accessible)
	e.GET("/login/google", receiver.GoogleOauthService.Login)
	e.GET("/login/google/start", receiver.GoogleOauthService.AuthUrl)
	e.GET("/login/refresh/:refresh_token", receiver.JwtService.RefreshToken)

	r := e.Group("/", receiver.JwtService.CreateMiddleware())

	r.GET("/projects", restricted)
	r.GET("/project/:id", restricted)

	r.POST("/project/save", restricted)
	r.POST("/project/build", restricted)

	e.Logger.Fatal(e.Start(":443"))

}
