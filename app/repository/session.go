package repository

import (
	"context"

	"github.com/emerfauzan/cake-store-api/app/model"
	"github.com/emerfauzan/cake-store-api/lib/logger"
)

func (repository *Repository) GetSessionBySessionKey(ctx context.Context, session_key string) (session model.Session, err error) {
	rows, err := repository.db.Query("select id, user_id, session_key, expires_at from sessions where session_key = (?)", session_key)

	if err != nil {
		logger.Error(ctx, "error get session by session key", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return session, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&session.Id, &session.UserId, &session.SessionKey, &session.ExpiresAt); err != nil {
			logger.Error(ctx, "error get session by session key", map[string]interface{}{
				"error": err,
				"tags":  []string{"query"},
			})
			return session, err
		}
	}

	return session, nil
}

func (repository *Repository) CreateSession(ctx context.Context, session model.Session) error {
	_, err := repository.db.Exec("INSERT INTO sessions (user_id, session_key, expires_at) values (?,?,?)", session.UserId, session.SessionKey, session.ExpiresAt)

	if err != nil {
		logger.Error(ctx, "error create session", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return err
	}
	return nil
}

func (repository *Repository) UpdateSession(ctx context.Context, session model.Session) error {
	_, err := repository.db.Exec("UPDATE sessions set user_id=?, session_key=?, expires_at=? WHERE id=?", session.UserId, session.SessionKey, session.ExpiresAt, session.Id)

	if err != nil {
		logger.Error(ctx, "error update session", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return err
	}
	return nil
}
