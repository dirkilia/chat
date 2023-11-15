package main

import (
	server "chat/internal/websocket"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	srv := server.New()
	http.Handle("/", http.FileServer(http.Dir("../templates")))
	http.Handle("/ws", websocket.Handler(srv.HandleWS))

	http.ListenAndServe(":8080", nil)
}
