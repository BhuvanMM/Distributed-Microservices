package main

import (
	"fmt"
	"log"
	"net/http"
)

const WebPort = "8080"

type Application struct {
}

func main() {
	log.Printf("Starting broker service on port %s\n", WebPort)
	app := Application{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.Routes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(fmt.Errorf("error starting http server: %v", err))
	}
}
