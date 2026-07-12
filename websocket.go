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
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}
////////////////////



func SendTime(conn *websocket.Conn ){
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


}


func ReadClient(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			return
		}
		log.Printf("Client said: %s", message)
	}
}


func ReadChat(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			return
		}
		log.Printf("CHAT SAID : %s", message)


		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
   			 log.Println("Write failed:", err)
   		 return
}


	}
}


func ChatHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	defer conn.Close()

	log.Println("Chat endpoint hit")

	ReadChat(conn)
}



func main() {


http.HandleFunc("/chats", ChatHandler)
   http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

go SendTime(conn)
ReadClient(conn)

}) 




http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "pages/wtest.html")
})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
