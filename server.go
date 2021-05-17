package main

import (
	"isar-golang-training-beginner/db"
	"isar-golang-training-beginner/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":80"))
}
