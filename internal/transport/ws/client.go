package ws

import (
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
