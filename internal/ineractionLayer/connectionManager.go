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

	broadcast chan InteractionEvent

	register chan *Client

	unregister chan *Client
}

func (cm *ConnectionManager) run() {
	for {
		select {
		case client := <-cm.register:
			cm.clients[client] = true
		case client := <-cm.unregister:
			if _, exists := cm.clients[client]; exists {
				delete(cm.clients, client)
				//do something to close channel
			}
		}
	}
}

func (cm *ConnectionManager) registerClient(context *gin.Context) {
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := Client{cm: cm, conn: conn}
	go client.receive()
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan InteractionEvent),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}
