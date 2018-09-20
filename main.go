package main

import (
	_ "github.com/joho/godotenv/autoload"

	"personnelMS-server/routers"
	"personnelMS-server/app/Models"
)

func main() {
	defer Models.CloseDB()
	router := routers.InitRouter()
	router.Run()
}