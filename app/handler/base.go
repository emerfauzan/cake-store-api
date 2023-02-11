package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emerfauzan/cake-store-api/app/usecase"
	"github.com/emerfauzan/cake-store-api/lib"
)

type Handler struct {
	usecase usecase.Usecase
}

type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Field   string `json:"field,omitempty"`
}

type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

type ResponseBody struct {
	Data    interface{}  `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
	Meta    ResponseMeta `json:"meta"`
}

type ResponseMeta struct {
	HTTPStatus int   `json:"http_status"`
	Total      *uint `json:"total,omitempty"`
	Offset     *uint `json:"offset,omitempty"`
	Limit      *uint `json:"limit,omitempty"`
	Page       *uint `json:"page,omitempty"`
}

func Init(usecase usecase.Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

func (handler *Handler) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}

func writeError(w http.ResponseWriter, err error) {
	var resp interface{}
	code := http.StatusInternalServerError

	fmt.Println(err)

	switch errOrig := err.(type) {
	case lib.CustomError:
		resp = ErrorBody{
			Errors: []ErrorInfo{
				{
					Message: errOrig.Message,
					Code:    errOrig.Code,
					Field:   errOrig.Field,
				},
			},
			Meta: ResponseMeta{
				HTTPStatus: errOrig.HTTPCode,
			},
		}

		code = errOrig.HTTPCode
	default:
		resp = ResponseBody{
			Message: "Internal Server Error",
			Meta: ResponseMeta{
				HTTPStatus: code,
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func writeSuccess(w http.ResponseWriter, data interface{}, message string, meta ResponseMeta) {
	resp := ResponseBody{
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.HTTPStatus)
	json.NewEncoder(w).Encode(resp)
}

func writeResponse(w http.ResponseWriter, resp interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
