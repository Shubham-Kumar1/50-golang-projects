package websocket

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	done       chan struct{}
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		done:       make(chan struct{}),
	}
}

func (pool *Pool) Start() {
	go func() {
		for {
			select {
			case client := <-pool.Register:
				pool.Clients[client] = true

				// Notify all clients about the new connection
				for c := range pool.Clients {
					c.Conn.WriteJSON(Message{
						Type: 1,
						Body: "New User Added",
					})
				}

			case client := <-pool.Unregister:
				delete(pool.Clients, client)

				// Notify remaining clients about the disconnection
				for c := range pool.Clients {
					c.Conn.WriteJSON(Message{
						Type: 1,
						Body: "One User Disconnected",
					})
				}

			case msg := <-pool.Broadcast:
				// Send the message to all clients including the sender
				for client := range pool.Clients {
					client.Conn.WriteJSON(msg)
				}

			case <-pool.done:
				return
			}
		}
	}()
}

func (pool *Pool) Stop() {
	close(pool.done)
}
