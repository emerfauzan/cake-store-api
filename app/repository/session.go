package repository

import (
	"github.com/emerfauzan/cake-store-api/app/model"
)

func (repository *Repository) CreateSession(session model.Session) error {
	_, err := repository.db.Exec("INSERT INTO sessions (user_id, session_key, expires_at) values (?,?,?)", session.UserId, session.SessionKey, session.ExpiresAt)

	if err != nil {
		return err
	}
	return nil
}
