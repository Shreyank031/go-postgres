package main

import (
	"fmt"
	"github.com/Shreyank031/go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
