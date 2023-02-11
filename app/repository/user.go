package repository

import (
	"context"

	"github.com/emerfauzan/cake-store-api/app/model"
	"github.com/emerfauzan/cake-store-api/lib/logger"
)

func (repository *Repository) GetUserByUsername(ctx context.Context, usename string) (user model.User, err error) {
	rows, err := repository.db.Query("select id, name, username, encrypted_password from users where username = (?)", usename)

	if err != nil {
		logger.Error(ctx, "error get user by username", map[string]interface{}{
			"error": err,
			"tags":  []string{"query"},
		})
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.EncryptedPassword); err != nil {
			return user, err
		}
	}

	return user, nil
}
