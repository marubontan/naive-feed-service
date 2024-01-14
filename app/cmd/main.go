package main

import (
	"naive-feed-service/app/server"
)

func main() {
	ginServer := server.NewServer()
	ginServer.Run()
}
