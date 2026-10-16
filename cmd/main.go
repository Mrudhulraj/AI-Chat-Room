package main

import (
	"live-video-meet/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
