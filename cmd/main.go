package main

import (
	server "gw-backend/internal"
	envVar "gw-backend/internal/envVars"
)

func main() {

	envVar.Start()

	server.Run()
}
