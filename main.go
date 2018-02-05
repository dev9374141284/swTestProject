package main

import (
	"log"
	"net/http"
	"TestProject/Route"
	"TestProject/Settings"
)

func main() {
	router := Route.NewRouter()
	log.Fatal(http.ListenAndServe(Settings.ServiceAddr, router))
}