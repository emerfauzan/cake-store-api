package main

import (
	"github.com/emerfauzan/cake-store-api/cli"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	cli.Execute()
}
