package main

import (
	"fmt"
	"net/http"

	"github.com/Serenity0204/react-golang-chat-app/pkg/websocket"
)

// set up websocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

	fmt.Println("WebSocket Endpoint Hit")

	// updrade connection to websocket
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Println(err)
		return
	}

	// create a new client everytime there is new connection then register it to the pool
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	// register for new client
	pool.Register <- client
	// Read the message from this client indefinitely
	client.Read()

}

func setupRoutes() {
	// create new pool
	pool := websocket.NewPool()

	// run pool.Start() whenever there's one case ready for the select statement
	go pool.Start()
	// whenever a user hit /ws, launch serveWs function
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

}

// run the app at port 8080
func main() {
	fmt.Println("App is running...")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
