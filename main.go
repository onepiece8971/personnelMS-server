package main

import (
	_ "github.com/joho/godotenv/autoload"

	"test/routers"
	"test/app/Models"
)

func main() {
	defer Models.CloseDB()
	router := routers.InitRouter()
	router.Run()
}