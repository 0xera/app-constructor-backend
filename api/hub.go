package api

import (
	"app-constructor-backend/model"
	"fmt"
	"log"
)

type SocketHub struct {
	subUser       string
	socketService *SocketService
	clients       map[*Client]bool
	register      chan *Client
	unregister    chan *Client
	quit          chan int
}

func NewSocketHub(socketService *SocketService, subUser string) *SocketHub {
	return &SocketHub{
		subUser:       subUser,
		socketService: socketService,
		clients:       make(map[*Client]bool),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		quit:          make(chan int),
	}
}

func (s *SocketHub) Run() {
	for {
		select {
		case client := <-s.register:
			s.onConnect(client)
		case client := <-s.unregister:
			s.onDisconnect(client)
		case <-s.quit:
			return
		}
	}
}

func (s *SocketHub) broadcast(data []byte, ignore *Client) {
	for c := range s.clients {
		if c != ignore {
			c.outbound <- data
		}
	}
}

func (s *SocketHub) send(data []byte, client *Client) {
	client.outbound <- data
}

func (s *SocketHub) onConnect(client *Client) {
	log.Println("client connected: ", client.socket.RemoteAddr())
	s.clients[client] = true
	count, err := s.socketService.repository.GetWidgetsCount(s.subUser)
	if err != nil {
		fmt.Println(err)
	}

	projects, err := s.socketService.repository.GetProjects(s.subUser)

	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := model.Response{
		WidgetsCount: count,
		Projects:     projects,
	}.MarshalJSON()

	s.send(bytes, client)
}

func (s *SocketHub) onDisconnect(client *Client) {

	client.close()

	delete(s.clients, client)

	if len(s.clients) == 0 {

		s.socketService.removeHub(s.subUser)
	}
}

func (s *SocketHub) onMessage(data []byte, client *Client) {
	r := &model.Response{}
	fmt.Println(data)

	if err := r.UnmarshalJSON(data); err != nil {
		fmt.Println(err)
	}

	go s.socketService.repository.UpdateUserProjects(s.subUser, r)

	bytes, err := r.MarshalJSON()
	if err != nil {
		return
	}

	s.broadcast(bytes, client)

}
