package handler

import (
	"encoding/json"
	"net/http"

	"github.com/emerfauzan/cake-store-api/app/request"
	"github.com/emerfauzan/cake-store-api/lib"
	"github.com/go-playground/validator/v10"
)

func (handler *Handler) AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, lib.ErrorInvalidParameter)
		return
	}

	err = validator.New().Struct(payload)
	if err != nil {
		writeError(w, lib.ErrorInvalidParameter)
		return
	}

	response, err := handler.usecase.AuthLoginUsecase(r.Context(), payload)

	if err != nil {
		writeError(w, err)
		return
	}

	writeSuccess(w, response, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}
