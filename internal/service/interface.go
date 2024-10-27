package service

import (
	"context"
	"time"
)

type Token struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Service interface {
	Payment(ctx context.Context, amount uint64, purpose string) (string, error)
	Status(ctx context.Context, id string) ([]byte, error)
}
