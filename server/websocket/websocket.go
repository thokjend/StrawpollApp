package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Struct to manage connections
type WebSocketHub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan string
	Mutex     sync.Mutex
}

var Hub = WebSocketHub{
	Clients:   make(map[*websocket.Conn]bool),
	Broadcast: make(chan string),
}

// WebSocket Handler
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade failed:", err)
		return
	}

	// Add client to the hub
	Hub.Mutex.Lock()
	Hub.Clients[conn] = true
	Hub.Mutex.Unlock()

	log.Println("Client connected")

	// Keep the connection alive and listen for messages
	go func() {
		defer func() {
			Hub.Mutex.Lock()
			delete(Hub.Clients, conn)
			Hub.Mutex.Unlock()
			conn.Close()
			log.Println("Client disconnected")
		}()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}
		}
	}()
}


// Broadcast messages to all connected clients
func BroadcastMessage(message string) {
	Hub.Mutex.Lock()
	defer Hub.Mutex.Unlock()
	for client := range Hub.Clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error writing to WebSocket:", err)
			client.Close()
			delete(Hub.Clients, client)
		}
	}
}