package main

import (
	"Conditions/db"
	sw "Conditions/routes"
	"fmt"
	"log"
)

func main() {
	db.Init()

	fmt.Println("Hello World")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8080"))
}
