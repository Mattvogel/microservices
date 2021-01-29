package main

import (
	"Auth/db"
	sw "Auth/routes"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println(os.Getenv("DB_PORT"))
	fmt.Println(os.Getenv("DB_NAME"))
	db.Init()

	db.InitRedis("1")
	log.Printf("Auth-Server Started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8000"))
}
