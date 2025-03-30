package main

import (
	"log"
	"net/http"

	"github.com/ojb112007/chat-app/chat/websocket"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", websocket.HandleConnections)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
