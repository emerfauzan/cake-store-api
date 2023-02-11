package main

import (
	"github.com/emerfauzan/cake-store-api/cli"
	"github.com/emerfauzan/cake-store-api/config"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	db := config.NewDB()
	cli.Execute()
	defer db.Close()
}
