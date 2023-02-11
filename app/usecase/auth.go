package usecase

import (
	"time"

	"github.com/emerfauzan/cake-store-api/app/model"
	"github.com/emerfauzan/cake-store-api/app/request"
	"github.com/emerfauzan/cake-store-api/app/response"
	"github.com/emerfauzan/cake-store-api/lib"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (usecase Usecase) AuthLoginUsecase(request request.LoginRequest) (response response.LoginResponse, err error) {
	var user = model.User{}
	user, err = usecase.repository.GetUserByUsername("cake_shop")

	if err != nil {
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(request.Password))

	if err != nil {
		return response, lib.ErrorWrongUsrnamePassword
	}

	uuid := uuid.NewString()
	expiresAt := time.Now().Add(48 * time.Hour)

	err = usecase.repository.CreateSession(model.Session{
		UserId:     user.Id,
		SessionKey: uuid,
		ExpiresAt:  expiresAt,
	})

	if err != nil {
		return response, err
	}

	response.Id = user.Id
	response.Name = user.Name
	response.Username = user.Username
	response.Token = uuid

	return response, nil
}
