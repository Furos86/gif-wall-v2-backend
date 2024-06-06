package interactionLayer

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	connectionManager *ConnectionManager
	clientConnection  *websocket.Conn
}

func (client *Client) Receive() {
	defer func() {
		// remove client from system
		client.connectionManager.unregister <- client
		client.clientConnection.Close()
	}()

	for {

		var event *InteractionEvent = &InteractionEvent{}

		err := client.clientConnection.ReadJSON(&event)

		if err != nil {
			// throw error when it it is not an expected disconnect
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
				websocket.CloseNoStatusReceived,
			) {
				log.Printf("client error: %v", err)
			}
			break
		}

		client.connectionManager.broadcast <- event
	}
}

func (client *Client) Send(event *InteractionEvent) {
	client.clientConnection.WriteJSON(event)
}
