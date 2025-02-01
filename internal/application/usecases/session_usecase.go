package usecases

import (
	"context"
	"core/config"
	"core/internal/domain"
)

type SessionUsecase struct {
	sessionRepo domain.SesssionRepository
	cfg         *config.Config
}

func NewSessionUsecase(sessionRepo domain.SesssionRepository, cfg *config.Config) domain.SessionUsecase {
	return &SessionUsecase{
		sessionRepo: sessionRepo,
		cfg:         cfg,
	}
}

func (s *SessionUsecase) CreateSession(ctx context.Context, session *domain.Session, expire int) (string, error) {
	return s.sessionRepo.CreateSession(ctx, session, expire)
}

func (s *SessionUsecase) DeleteByID(ctx context.Context, sessionID string) error {
	return s.sessionRepo.DeleteByID(ctx, sessionID)
}

func (s *SessionUsecase) GetSessionByID(ctx context.Context, sessionID string) (*domain.Session, error) {
	return s.sessionRepo.GetSessionByID(ctx, sessionID)
}
