package repository

import (
	"context"
	"core/config"
	"core/internal/domain"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type SesssionRepository struct {
	redisClient *redis.Client
	basePrefix  string
	cfg         *config.Config
}

func NewSessionRepository(redisClient *redis.Client, cfg *config.Config) domain.SesssionRepository {
	return &SesssionRepository{
		redisClient: redisClient,
		basePrefix:  cfg.Session.Prefix,
		cfg:         cfg,
	}
}

func (r *SesssionRepository) CreateSession(ctx context.Context, session *domain.Session, expire int) (string, error) {
	session.SessionID = uuid.New().String()
	sessionKey := r.createKey(session.SessionID)

	sessBytes, err := json.Marshal(&session)
	if err != nil {
		return "", errors.WithMessage(err, "sessionRepo.CreateSession.json.Marshal")
	}
	if err = r.redisClient.Set(ctx, sessionKey, sessBytes, time.Second*time.Duration(expire)).Err(); err != nil {
		return "", errors.Wrap(err, "sessionRepo.CreateSession.redisClient.Set")
	}
	return sessionKey, nil
}

func (r *SesssionRepository) GetSessionByID(ctx context.Context, sessionID string) (*domain.Session, error) {
	sessBytes, err := r.redisClient.Get(ctx, sessionID).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "sessionRep.GetSessionByID.redisClient.Get")
	}

	sess := &domain.Session{}
	if err = json.Unmarshal(sessBytes, &sess); err != nil {
		return nil, errors.Wrap(err, "sessionRepo.GetSessionByID.json.Unmarshal")
	}
	return sess, nil
}

func (r *SesssionRepository) DeleteByID(ctx context.Context, sessionID string) error {
	if err := r.redisClient.Del(ctx, sessionID).Err(); err != nil {
		return errors.Wrap(err, "sessionRepo.DeleteByID")
	}
	return nil
}

func (s *SesssionRepository) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.basePrefix, sessionID)
}
