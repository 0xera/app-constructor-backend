package api

import (
	"github.com/gorilla/websocket"
	"net"
)

type Client struct {
	service  *SocketHub
	addr     net.Addr
	socket   *websocket.Conn
	outbound chan []byte
}

func newClient(hub *SocketHub, socket *websocket.Conn) *Client {
	return &Client{
		service:  hub,
		socket:   socket,
		addr:     socket.RemoteAddr(),
		outbound: make(chan []byte),
	}
}

func (client *Client) read() {
	defer func() {
		client.service.unregister <- client
	}()
	for {
		_, data, err := client.socket.ReadMessage()
		if err != nil {
			break
		}
		client.service.onMessage(data, client)
	}
}

func (client *Client) write() {
	for {
		select {
		case data, ok := <-client.outbound:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.socket.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (client Client) close() {
	client.socket.Close()
	close(client.outbound)
}
