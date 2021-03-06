package main

import (
	"Devices/db"
	sw "Devices/routes"
	"fmt"
	"log"
)

func main() {
	db.Init()

	fmt.Println("Hello World")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8080"))
}
