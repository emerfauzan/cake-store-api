package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emerfauzan/cake-store-api/app/request"
	"github.com/emerfauzan/cake-store-api/lib"
	"github.com/gorilla/mux"
)

func (handler *Handler) GetCakesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value("CurrentUserID").(uint)
	if !ok {
		writeError(w, lib.ErrorForbidden)
		return
	}

	result, err := handler.usecase.GetCakes(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}

	writeSuccess(w, result, "success", ResponseMeta{HTTPStatus: http.StatusOK})

}

func (handler *Handler) GetCakeByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value("CurrentUserID").(uint)
	if !ok {
		writeError(w, lib.ErrorForbidden)
		return
	}

	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	result, err := handler.usecase.GetCakeById(r.Context(), uint(id))
	if err != nil {
		writeError(w, err)
		return
	}

	writeSuccess(w, result, "success", ResponseMeta{HTTPStatus: http.StatusOK})

}

func (handler *Handler) CreateCakeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	CurrentUserID, ok := ctx.Value("CurrentUserID").(uint)
	if !ok {
		writeError(w, lib.ErrorForbidden)
		return
	}

	var payload request.CreateCakeRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, lib.ErrorInvalidParameter)
		return
	}

	payload.UserId = CurrentUserID

	err = handler.usecase.CreateCake(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}

	writeSuccess(w, nil, "success", ResponseMeta{HTTPStatus: http.StatusCreated})

}

func (handler *Handler) UpdateCakeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	CurrentUserID, ok := ctx.Value("CurrentUserID").(uint)
	if !ok {
		writeError(w, lib.ErrorForbidden)
		return
	}

	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	var payload request.UpdateCakeRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, lib.ErrorInvalidParameter)
		return
	}

	payload.UserId = CurrentUserID

	err = handler.usecase.UpdateCake(r.Context(), payload, uint(id))
	if err != nil {
		writeError(w, err)
		return
	}

	writeSuccess(w, nil, "success", ResponseMeta{HTTPStatus: http.StatusOK})

}

func (handler *Handler) DeleteCakeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	CurrentUserID, ok := ctx.Value("CurrentUserID").(uint)
	if !ok {
		writeError(w, lib.ErrorForbidden)
		return
	}

	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	err := handler.usecase.DeleteCake(r.Context(), uint(id), CurrentUserID)
	if err != nil {
		writeError(w, err)
		return
	}

	writeSuccess(w, nil, "success", ResponseMeta{HTTPStatus: http.StatusOK})

}
