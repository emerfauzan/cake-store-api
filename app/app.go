package app

import (
	"net/http"
	"time"

	"github.com/emerfauzan/cake-store-api/app/handler"
	"github.com/emerfauzan/cake-store-api/app/repository"
	"github.com/emerfauzan/cake-store-api/app/usecase"
	"github.com/emerfauzan/cake-store-api/config"
	"github.com/gorilla/mux"
)

func NewApp() {
	r := mux.NewRouter()
	db := config.NewDB()
	repository := repository.Init(db)
	usecase := usecase.Init(*repository)
	handler := handler.Init(*usecase)

	r.HandleFunc("/", handler.HealthzHandler)
	r.HandleFunc("/auth/login", handler.AuthLoginHandler).Methods("POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	srv.ListenAndServe()
}
