package api

import (
	"app-constructor-backend/auth"
	"app-constructor-backend/model"
	"app-constructor-backend/repository"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type SocketService struct {
	jwtService *auth.JwtService
	repository *repository.Repository
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func NewSocketService(repository *repository.Repository, jwtService *auth.JwtService) *SocketService {
	return &SocketService{
		repository: repository,
		jwtService: jwtService,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (s *SocketService) Run() {
	for {
		select {
		case client := <-s.register:
			s.onConnect(client)
		case client := <-s.unregister:
			s.onDisconnect(client)
		}
	}
}

func (s *SocketService) broadcast(data []byte, ignore *Client) {
	for c := range s.clients {
		if c != ignore {
			c.outbound <- data
		}
	}
}

func (s *SocketService) send(data []byte, client *Client) {
	client.outbound <- data
}

func (s *SocketService) onConnect(client *Client) {
	log.Println("client connected: ", client.socket.RemoteAddr())

	s.clients[client] = true

	s.send([]byte("for all"), client)
	s.broadcast([]byte(fmt.Sprintf("for all expect client with addr = %d", client.addr)), client)
}

func (s *SocketService) onDisconnect(client *Client) {
	log.Println("client disconnected: ", client.socket.RemoteAddr())

	client.close()

	delete(s.clients, client)

	s.broadcast([]byte(fmt.Sprintf("Client left with addr = %d", client.addr)), nil)
}

func (s *SocketService) onMessage(data []byte, client *Client) {
	r := &model.Response{}
	if err := r.UnmarshalJSON(data); err != nil {
		fmt.Println(err)
	}
	log.Println("onMessage: ", string(data))
	s.broadcast(data, client)
}

func (s *SocketService) Connect(c echo.Context) error {
	tokenReq := c.Param("token")
	fmt.Println(tokenReq)
	_, err := s.jwtService.ParseToken(tokenReq, os.Getenv("SECRET_JWT_SOCKET"))

	if err != nil || !s.jwtService.ContainsToken(tokenReq) {
		return c.JSON(http.StatusInternalServerError, "token parse error")
	}

	s.jwtService.DeleteToken(tokenReq)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Error(err)
		return err
	}

	client := newClient(s, ws)
	s.register <- client

	go client.read()
	go client.write()
	return nil
}
