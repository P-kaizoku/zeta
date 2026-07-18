package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByGoogleID(ctx context.Context, googleID string) (*User, error)
	Create(ctx context.Context, user *User) error
}

type RoomRepository interface {
	Create(ctx context.Context, room *Room) error
	GetByID(ctx context.Context, id uuid.UUID) (*Room, error)
	ListForUser(ctx context.Context, userID uuid.UUID) ([]*Room, error)
	AddMember(ctx context.Context, roomID, userID uuid.UUID) error
	IsMember(ctx context.Context, roomID, userID uuid.UUID) (bool, error)
}

type MessageRepository interface {
	Create(ctx context.Context, message *Message) error
	ListRecent(ctx context.Context, roomId uuid.UUID, limit int, beforeTime time.Time) ([]*Message, error)
}
