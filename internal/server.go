package server

import (
	"github.com/gin-gonic/gin"
	"gw-backend/internal/envVars"
	"gw-backend/internal/services/interactionLayer"
)

func Run() {

	if envVar.Variables.Environment == "production" {
		gin.SetMode("release")
	}

	router := gin.Default()

	setTrustedProxiesError := router.SetTrustedProxies(nil)

	if setTrustedProxiesError != nil {
		panic(setTrustedProxiesError)
	}

	connectionManager := interactionLayer.NewConnectionManager()

	go connectionManager.Run()

	router.GET("/ws", func(context *gin.Context) {
		print("we got connection")
		connectionManager.RegisterClient(context)

	})

	runError := router.Run("127.0.0.1:" + envVar.Variables.Port)

	if runError != nil {
		panic(runError)
	}
}
