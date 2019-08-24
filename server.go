package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type HelloServer struct {
	port int
}

func NewHelloServer(port int) *HelloServer {
	return &HelloServer{port: port}
}

func (server HelloServer) Start() {
	http.HandleFunc("/hello", server.HelloServerHandler)
	http.ListenAndServe(":"+strconv.Itoa(server.port), nil)
}

func (server HelloServer) HelloServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "REAL WORLD")
}
