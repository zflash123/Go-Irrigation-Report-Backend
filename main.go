package main

import (
	"go-irrigation-report-api/routes"
)

func runServer() {
	routes.Routes()
}

func main() {
	runServer()
}
