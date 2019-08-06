package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	http.HandleFunc("/ws", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = originChecker
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected")

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func reader(con *websocket.Conn) {
	for {
		messageType, p, err := con.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := con.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func originChecker(r *http.Request) bool {
	return true
}
