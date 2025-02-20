package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", msg)
		// Echo the message back (we'll improve this later)
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", handleWebSocket)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}