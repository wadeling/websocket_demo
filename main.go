package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var server *WsServer

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域连接
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Received message: %s,message type:%v\n", p, messageType)

		rspMsg := fmt.Sprintf("server recieved:%s", p)
		err = conn.WriteMessage(messageType, []byte(rspMsg))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func handleConnectionsV2(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = server.AddConn(conn)
	fmt.Println("new client connected,add to pool")
	return

}

func main() {
	// create conn pool
	server = NewWsServer()
	go func() {
		_ = server.Run()
	}()

	// create server
	http.HandleFunc("/ws", handleConnectionsV2)
	fmt.Println("WebSocket server started")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
