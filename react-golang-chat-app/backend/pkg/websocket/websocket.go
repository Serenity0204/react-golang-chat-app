package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// define upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// upgrade the connection to web socket
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}
