package main

import (
	"cobalagi/db"
	"cobalagi/routes"
	"os"
)

func main() {

	db.Init()
	route := routes.Init()
	port := os.Getenv("PORT")
	route.Logger.Fatal(route.Start(":" + port))
}
