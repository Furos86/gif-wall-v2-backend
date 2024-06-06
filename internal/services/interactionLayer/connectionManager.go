package interactionLayer

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ConnectionManager struct {
	// Connected clients
	clients map[*Client]bool

	broadcast chan *InteractionEvent

	unregister chan *Client
}

func (cm *ConnectionManager) Run() {
	// main connection manager loop
	for {
		select {
		case client := <-cm.unregister:
			delete(cm.clients, client)

		case event := <-cm.broadcast:
			for client := range cm.clients {
				client.Send(event)
			}
		}
	}
}

func (cm *ConnectionManager) RegisterClient(context *gin.Context) {
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{connectionManager: cm, clientConnection: conn}

	cm.clients[client] = true

	go client.Receive()
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *InteractionEvent),
		unregister: make(chan *Client),
	}
}
