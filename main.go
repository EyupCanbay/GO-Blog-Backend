package main

import (
	"fmt"
	"log"
	"net/http"

	router "golang-blog-backend/router"
)

func main() {
	r := router.NewRouter()

	fmt.Println("server is running at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
