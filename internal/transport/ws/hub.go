package ws

import (
	"encoding/json"

	"github.com/P-kaizoku/zeta/internal/domain"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *domain.Message
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			payload, err := json.Marshal(message)
			if err != nil {
				continue
			}

			for client := range h.clients {
				if client.roomIDs[message.RoomID] {
					select {
					case client.send <- payload:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}
