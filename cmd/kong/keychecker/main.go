package main

import (
	"github.com/Kong/go-pdk/server"
	kongdelivery "github.com/jiramot/demo-kong-plugin/internal/delivery/kong"
)

func main() {
	server.StartServer(kongdelivery.New, "0.1", 0)
}