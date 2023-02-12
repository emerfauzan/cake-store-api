package test

import (
	"github.com/emerfauzan/cake-store-api/app/handler"
	"github.com/emerfauzan/cake-store-api/app/repository"
	"github.com/emerfauzan/cake-store-api/app/usecase"
	"github.com/emerfauzan/cake-store-api/config"
)

type Test struct {
	handler *handler.Handler
	token   string
}

func Init() Test {
	db := config.NewDB()
	repository := repository.Init(db)
	usecase := usecase.Init(*repository)
	handler := handler.Init(*usecase)
	return Test{
		handler: &handler,
	}
}
