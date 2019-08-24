package main

import (
	"log"
)

func main () {
	config, err := NewConfig("real.json")

	if err != nil {
		panic(err)
	}
	
	server := NewHelloServer(config.Port)

	log.Println("Handling server on port", config.Port)

	server.Start()
}