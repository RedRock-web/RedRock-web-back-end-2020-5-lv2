package main

import (
	"RedRock-web-back-end-2020-5-lv2/app"
	"RedRock-web-back-end-2020-5-lv2/database"
	"RedRock-web-back-end-2020-5-lv2/router"
)

func main() {
	database.ConnetDb()
	database.CreateTable()
	app.GetAllStudentsInfo()
	router.SetupRouter()
}

