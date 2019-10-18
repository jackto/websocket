package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients sync.Map
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/count", countUser)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients.Store(ws, true)
	for {
		var msg Message
		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error: %v", err)
			clients.Delete(ws)
			//delete(clients,ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		clients.Range(func(key, value interface{}) bool {
			switch client := key.(type) {
			case *websocket.Conn:
				err := client.WriteJSON(msg)
				if err != nil {
					fmt.Println("err %s", err)
					client.Close()
					clients.Delete(client)
				}
			}
			return true
		})
	}
}

func countUser(w http.ResponseWriter, r *http.Request) {

	var i = 0
	clients.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	fmt.Fprintf(w, "%d", i)
}
