package domain

import (
	"context"

	"github.com/google/uuid"
)

type Session struct {
	SessionID string    `json:"session_id" redis:"session_id"`
	UserID    uuid.UUID `json:"user_id" redis:"user_id"`
}

type SesssionRepository interface {
	CreateSession(ctx context.Context, session *Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}

type SessionUsecase interface {
	CreateSession(ctx context.Context, session *Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
