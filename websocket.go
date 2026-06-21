package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)


///////shift http -> ws ////////

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
////////////////////

func main() {



   http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Send time every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now().Format("15:04:05")
		err := conn.WriteMessage(websocket.TextMessage, []byte(now))
		if err != nil {
			log.Println("Client disconnected:", err)
			return
		}
	}
}) 




http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "pages/wtest.html")
})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
