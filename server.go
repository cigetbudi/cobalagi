package main

import (
	"cobalagi/db"
	"cobalagi/routes"
)

func main() {

	db.Init()
	route := routes.Init()
	route.Logger.Fatal(route.Start(":8080"))
}
