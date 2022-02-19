package websocket

import "fmt"

type Pool struct {
	Register   chan *Client     // show new user joined when new user is connected
	Unregister chan *Client     // show user disconnected when a user is dc
	Clients    map[*Client]bool // a map of users with boolean value indicating active/inactive user
	Broadcast  chan Message     // a channel of message that will loop through all the clients and send them a message through socket connection
}

// create new pool
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		// user register
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// tell each client in the pool, new user joined
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Has Joined"})
			}
			break
		// case 2
		case client := <-pool.Unregister:
			// delete disconnected client
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// tell each client in the pool, one user disconnected
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Has Disconnected"})
			}
			break
		// case 3
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			// broadcasting message to each of the client
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
