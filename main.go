package main

import (
	"github.com/emerfauzan/cake-store-api/cli"
	"github.com/emerfauzan/cake-store-api/lib/logger"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	logger.Init()
}

func main() {
	cli.Execute()
}
