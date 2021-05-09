package api

import (
	"app-constructor-backend/auth"
	"app-constructor-backend/model"
	"app-constructor-backend/repository"
	"app-constructor-backend/task"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	hubMap = make(map[string]*SocketHub)
)

type SocketService struct {
	jwtService  *auth.JwtService
	repository  *repository.Repository
	taskService *task.Service
}

func NewSocketService(taskService *task.Service, repository *repository.Repository, jwtService *auth.JwtService) *SocketService {
	return &SocketService{
		taskService: taskService,
		repository:  repository,
		jwtService:  jwtService,
	}
}

func (s *SocketService) ConnectToCollaborate(c echo.Context) error {
	tokenReq := c.Param("token")
	fmt.Println(tokenReq)
	token, err := s.jwtService.ParseToken(tokenReq, os.Getenv("SECRET_JWT_SOCKET"))

	if err != nil || !s.jwtService.ContainsToken(tokenReq) {
		return c.JSON(http.StatusInternalServerError, "token parse error")
	}

	s.jwtService.DeleteToken(tokenReq)

	userClaims, ok := token.Claims.(*model.UserClaims)

	if ok && token.Valid {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			c.Error(err)
			return err
		}

		hub, ok := hubMap[userClaims.Sub]
		if !ok {
			hub = NewSocketHub(s, userClaims.Sub)
			hubMap[userClaims.Sub] = hub
			go hub.Run()
		}
		client := newClient(hub, ws)
		hub.register <- client

		go client.read()
		go client.write()
	}

	return nil
}

func (s *SocketService) ConnectToBuild(c echo.Context) error {
	tokenReq := c.Param("token")
	fmt.Println(tokenReq)
	token, err := s.jwtService.ParseToken(tokenReq, os.Getenv("SECRET_JWT_SOCKET"))

	if err != nil || !s.jwtService.ContainsToken(tokenReq) {
		return c.JSON(http.StatusInternalServerError, "token parse error")
	}

	s.jwtService.DeleteToken(tokenReq)

	userClaims, ok := token.Claims.(*model.UserClaims)
	projectIdString := c.QueryParam("id")
	if ok && token.Valid && projectIdString != "" {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			c.Error(err)
			return err
		}

		projectId, err := strconv.Atoi(projectIdString)
		if err != nil {
			err := closeSocket(c, ws)
			return err
		}
		filename, err := s.taskService.Execute(userClaims.Sub, projectId)
		if err != nil {
			_ = closeSocket(c, ws)
			return err
		}
		if filename == "" {
			_ = closeSocket(c, ws)
			return err
		}
		err = ws.WriteMessage(websocket.TextMessage, []byte(filename))
		if err != nil {
			c.Logger().Error(err)
		}
		_ = closeSocket(c, ws)
	}

	return nil
}

func closeSocket(c echo.Context, ws *websocket.Conn) error {
	err := ws.WriteMessage(websocket.CloseMessage, []byte{})
	if err != nil {
		c.Logger().Error(err)
	}
	return err
}

func (s *SocketService) removeHub(hubKey string) {
	delete(hubMap, hubKey)
}
