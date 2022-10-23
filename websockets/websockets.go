package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		if err = websocket.Message.Recieve(ws, &reply); err != nil {
			fmt.Println("Cant Recieve")
			break
		}
		fmt.Println("Recieved back form client: " + reply)
		msg := "Recieved: " + reply
		fmt.Println("Sending to client:" + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {

			fmt.Println("Cnat send")
			break
		}
	}
}
func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Listenandserveerror:", err)
	}
}
