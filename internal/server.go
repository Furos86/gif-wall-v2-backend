package server

import (
	envVar "gw-backend/internal/envVars"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Run() {

	if envVar.Variables.Environment == "production" {
		gin.SetMode("release")
	}

	router := gin.Default()

	setTrustedProxiesError := router.SetTrustedProxies(nil)

	if setTrustedProxiesError != nil {
		panic(setTrustedProxiesError)
	}

	connectionManager = NewConnectionManager()

	router.GET("/ws", func(context *gin.Context) {
		print("we got connection")
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

	})

	runError := router.Run("127.0.0.1:" + envVar.Variables.Port)

	if runError != nil {
		panic(runError)
	}
}
