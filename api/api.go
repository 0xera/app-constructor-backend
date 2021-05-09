package api

import (
	"app-constructor-backend/auth"
	"app-constructor-backend/repository"
	"app-constructor-backend/task"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

type Service struct {
	Repository         *repository.Repository
	JwtService         *auth.JwtService
	GoogleOauthService *auth.GoogleOauthService
	SocketService      *SocketService
	TaskService        *task.Service
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
	e.GET("/login/refresh/:refreshToken", service.JwtService.RefreshToken)

	r := e.Group("/main", service.JwtService.CreateMiddleware())

	r.GET("/ws/token", service.JwtService.SocketToken)

	//r.GET("/projects", service.Repository.GetProjects)
	//r.POST("/project/save", service.Repository.SaveProject)
	//r.POST("/project/delete", service.Repository.DeleteProject)
	r.POST("/project/create", service.Repository.CreateProject)
	r.GET("/project/download/:name", service.Repository.DownloadProject)

	r.POST("/project/build", service.Repository.Restricted)

	e.GET("ws/:token", service.SocketService.ConnectToCollaborate)
	e.GET("ws/build/:token", service.SocketService.ConnectToBuild)

	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("SERVER_PORT")))

}
