package repository

import (
	"github.com/emerfauzan/cake-store-api/app/model"
)

func (repository *Repository) GetUserByUsername(usename string) (user model.User, err error) {
	// var result = model.User{}
	rows, err := repository.db.Query("select id, name, username, encrypted_password from users where username = (?)", usename)

	if err != nil {
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
