package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string          // unique id to each connection
	Conn *websocket.Conn // pointer to websocket.Conn
	Pool *Pool           // pointer to Pool which contains all of the client
}

// Message struct
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// The client read method, that will be listening to new coming message from this client
func (c *Client) Read() {
	// handle when quitting
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// read new message
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// construct new message
		message := Message{Type: messageType, Body: string(p)}

		// if any new message, it will be passed through the Pool.Broadcast, which will send message to each of the client in the pool
		c.Pool.Broadcast <- message

		fmt.Println(message)

	}
}
