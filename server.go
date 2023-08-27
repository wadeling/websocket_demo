package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// WsServer websocket server
type WsServer struct {
	pool *ConnPool
}

func (w *WsServer) Run() error {
	num := 0
	for {
		time.Sleep(5 * time.Second)
		num++
		err := w.pool.PublishTextMsg(fmt.Sprintf("publish msg %v", num))
		if err != nil {
			log.Printf("publish err:%v \n", err)
		} else {
			log.Println("publish msg ok")
		}
	}
}

func (w *WsServer) AddConn(conn *websocket.Conn) error {
	w.pool.AddConn(conn)
	return nil
}

func NewWsServer() *WsServer {
	w := &WsServer{}
	w.pool = NewConnPool()
	return w
}
