package interactionLayer

import (
	"encoding/json"
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

		// check if the error is a websocket close error
		if err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
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

			if _, ok := err.(*json.SyntaxError); ok {
				log.Printf("JSON syntax error: %v", err)
			} else {
				log.Printf("error: %v", err)
			}
		} else {
			client.connectionManager.broadcast <- event
		}
	}
}

func (client *Client) Send(event *InteractionEvent) {
	client.clientConnection.WriteJSON(event)
}
