package main

import (
	"go-irrigation-report-backend/routes"
	"go-irrigation-report-backend/config"
)

func runServer() {
	config.ViperEnvConfig()
	routes.Routes()
}

func main() {
	runServer()
}
