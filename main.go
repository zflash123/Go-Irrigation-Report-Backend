package main

import (
	"go-irrigation-report-backend/routes"
)

func runServer() {
	routes.Routes()
}

func main() {
	runServer()
}
