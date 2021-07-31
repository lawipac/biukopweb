package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
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

		switch messageType {
		case websocket.TextMessage:
			WsProcessingTxtMessage(conn, string(p))
			break
		case websocket.BinaryMessage:
			WsProcessingBinaryMessage(conn, p)
			break
		case websocket.PingMessage:
			break
		case websocket.PongMessage:
			break
		}
	}
}

func WsProcessingTxtMessage(conn *websocket.Conn, str string) {
	if str == "send dummy string for 500 times" {
		go wsDummySender(conn)
	} else {
		WsEchoIncomingMessage(conn, str, websocket.TextMessage)
	}

}

func WsProcessingBinaryMessage(conn *websocket.Conn, data []byte) {
	return
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

func wsDummySender(conn *websocket.Conn) {
	//write subsequent 5 copies, each after 1 second
	log.Info("start sending server data to client ..")
	p := "From server "
	for i := 1; i < 500; i++ {
		time.Sleep(1 * time.Second)
		n := time.Now()
		s := n.Format("2006-01-02 15:04:05")
		p1 := p + s
		//msg := fmt.Sprintf("copy %d, %s, %s", i, p, wsDummyString()) // 4M long string no issue
		msg := fmt.Sprintf("copy %d, %s ", i, p1)
		log.Info("dummy sender is working, ", msg)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Println("wsDummySender stopped on error: ", err)
			return
		}
	}
}
