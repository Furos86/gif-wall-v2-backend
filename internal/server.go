package server

import (
	envVar "gw-backend/internal/envVars"

	"github.com/gin-gonic/gin"
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

	router.GET("/ws", func (context *gin.Context) {
		print("we got connection")
	})

	runError := router.Run("127.0.0.1:" + envVar.Variables.Port)

	if runError != nil {
		panic(runError)
	}
}