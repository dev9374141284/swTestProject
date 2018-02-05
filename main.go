package main

import (
	"log"
	"net/http"
	"TestProject/Route"
	"TestProject/Settings"
	"TestProject/Business"
)

func main() {
	router := Route.NewRouter()
	log.Fatal(http.ListenAndServe(Settings.ServiceAddr, router))
	defer Business.CloseSession()
}