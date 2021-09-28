package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/websocket"
)

func basic_route(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Simple Server")
}

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a websocket
	// connection

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func main() {
	fmt.Println("Chat App v0.01")
	http.HandleFunc("/", basic_route)
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8080", nil)
}
