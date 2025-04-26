package websocket

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type      int       `json:"type"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
	ClientID  string    `json:"client_id,omitempty"`
}

func NewClient(conn *websocket.Conn, pool *Pool) *Client {
	return &Client{
		ID:   uuid.New().String(),
		Conn: conn,
		Pool: pool,
	}
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}

		message := Message{
			Type:      msgType,
			Body:      string(msg),
			Timestamp: time.Now(),
			ClientID:  c.ID,
		}

		c.Pool.Broadcast <- message
	}
}
