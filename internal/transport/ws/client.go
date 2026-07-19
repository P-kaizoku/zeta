package ws

import (
	"time"

	"github.com/P-kaizoku/zeta/internal/domain"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	send    chan []byte
	userID  uuid.UUID
	roomIDs map[uuid.UUID]bool
}

func (c *Client) writePump() {
	defer c.conn.Close()
	for message := range c.send {
		c.conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (c *Client) readPump() {
	defer c.conn.Close()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			c.hub.unregister <- c
			break
		}
		payload := &domain.Message{
			ID:        uuid.New(),
			RoomID:    c.roomID,
			SenderID:  c.userID,
			Content:   string(msg),
			CreatedAt: time.Now(),
		}
		c.hub.broadcast <- payload

	}
}
