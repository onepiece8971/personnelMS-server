package main

import (
	_ "github.com/joho/godotenv/autoload"

	"trunk/develop/routers"
	"trunk/develop/app/Models"
)

func main() {
	defer Models.CloseDB()
	router := routers.InitRouter()
	router.Run()
}