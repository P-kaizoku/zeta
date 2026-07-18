package main

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	Name      string
	IsDM      bool
	CreatedAt time.Time
}

type RoomMember struct {
	RoomID   uuid.UUID
	UserID   uuid.UUID
	JoinedAt time.Time
}
