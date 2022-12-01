package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kwamekyeimonies/stocks-api/database"
	"github.com/kwamekyeimonies/stocks-api/router"
)

func main() {
	r := router.Router()
	port := ":8090"

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db := database.Create_connection()

	db.Close()

	fmt.Println("Server running on port ", port)

	log.Fatal(http.ListenAndServe(port, r))
}
