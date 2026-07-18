package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	GoogleID    string
	Email       string
	DisplayName string
	AvatarURL   string
	CreatedAt   time.Time
}
