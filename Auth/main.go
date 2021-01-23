package main

import (
	"Auth/db"
	sw "Auth/routes"
	"log"
)

func main() {

	db.Init()

	db.InitRedis("1")
	log.Printf("Auth-Server Started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8000"))
}
