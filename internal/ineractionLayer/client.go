package interactionLayer

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	cm   *ConnectionManager
	conn *websocket.Conn
}

func (client *Client) receive() {
	for {

		event := &InteractionEvent{}

		err := client.conn.ReadJSON(*event)

		if err != nil {
			log.Println(err)
		}

		client.cm.broadcast <- *event
	}
}
