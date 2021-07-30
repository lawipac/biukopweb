package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const apiV1WebSocket = apiV1Prefix + "ws"

// We'll need to define an upgradeToWs
// this will require a Read and Write buffer size
var upgradeToWs = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func apiV1WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgradeToWs.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgradeToWs.Upgrade(w, r, nil)
	if err != nil {
		log.Println("cannot upgrade websocket", err)
		return
	}

	// helpful log statement to show connections
	log.Println("Websocket Api/V1: Client Connected", r.RemoteAddr)

	wsReader(ws)
}

func wsReader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		WsEchoIncomingMessage(conn, string(p), messageType)
		//switch messageType {
		//case websocket.TextMessage:
		//	WsProcessingTxtMessage(conn, string(p))
		//	break
		//case websocket.BinaryMessage:
		//	WsProcessingBinaryMessage(conn, p)
		//	break
		//case websocket.PingMessage:
		//	break
		//case websocket.PongMessage:
		//	break
		//}
	}
}

func WsEchoIncomingMessage(conn *websocket.Conn, msg string, messageType int) {
	// print out that message for clarity
	fmt.Println(msg)

	// this is echo
	if err := conn.WriteMessage(messageType, []byte(msg)); err != nil {
		log.Println(err)
		return
	}
}
