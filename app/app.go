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

	r.HandleFunc("/", handler.StandardMiddleware(http.HandlerFunc(handler.HealthzHandler)).ServeHTTP).Methods("GET")
	r.HandleFunc("/auth/login", handler.StandardMiddleware(http.HandlerFunc(handler.AuthLoginHandler)).ServeHTTP).Methods("POST")
	r.HandleFunc("/cakes", handler.UserMiddlewares(http.HandlerFunc(handler.GetCakesHandler)).ServeHTTP).Methods("GET")
	r.HandleFunc("/cakes/{id:[0-9]+}", handler.UserMiddlewares(http.HandlerFunc(handler.GetCakeByIdHandler)).ServeHTTP).Methods("GET")
	r.HandleFunc("/cakes", handler.UserMiddlewares(http.HandlerFunc(handler.CreateCakeHandler)).ServeHTTP).Methods("POST")
	r.HandleFunc("/cakes/{id:[0-9]+}", handler.UserMiddlewares(http.HandlerFunc(handler.UpdateCakeHandler)).ServeHTTP).Methods("PUT")
	r.HandleFunc("/cakes/{id:[0-9]+}", handler.UserMiddlewares(http.HandlerFunc(handler.DeleteCakeHandler)).ServeHTTP).Methods("DELETE")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	srv.ListenAndServe()
}
