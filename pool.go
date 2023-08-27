package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type ConnPool struct {
	sync.RWMutex
	connections []*websocket.Conn
}

func (c *ConnPool) AddConn(conn *websocket.Conn) {
	c.Lock()
	defer c.Unlock()
	c.connections = append(c.connections, conn)
}

func (c *ConnPool) PublishTextMsg(msg string) error {
	c.Lock()
	defer c.Unlock()
	for _, v := range c.connections {
		err := v.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Printf("failed to publish msg to client:%v\n", v.LocalAddr())
			return err
		}
	}
	return nil
}

func NewConnPool() *ConnPool {
	c := &ConnPool{}
	return c
}
