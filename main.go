package main

import (
	"fmt"
	"go-notes-api/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	router.RegisterRoutes(r)

	fmt.Println("8085 portunda sunucu calisiyor")
	log.Fatal(http.ListenAndServe(":8085", r))
}
