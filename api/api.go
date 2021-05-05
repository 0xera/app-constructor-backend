package api

import (
	"app-constructor-backend/auth"
	"app-constructor-backend/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

type Service struct {
	Repository         *repository.Repository
	JwtService         *auth.JwtService
	GoogleOauthService *auth.GoogleOauthService
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Hello :)")
}

func (service *Service) Serve() {

	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", accessible)
	e.GET(os.Getenv("REDIRECT_PATH"), service.GoogleOauthService.Login)
	e.GET("/login/google/start", service.GoogleOauthService.AuthUrl)
	e.GET("/login/refresh/:refresh_token", service.JwtService.RefreshToken)

	r := e.Group("/main", service.JwtService.CreateMiddleware())

	r.GET("/projects", service.Repository.GetProjects)

	r.POST("/project/save", service.Repository.SaveProject)
	r.POST("/project/delete", service.Repository.DeleteProject)
	r.POST("/project/create", service.Repository.CreateProject)

	r.POST("/project/build", service.Repository.Restricted)

	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("SERVER_PORT")))

}
