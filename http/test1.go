package main

import (
	"fmt"
	"net/http"
	// "github.com/samarth-nagar/go-apis/weather"
	"log"
	"time"
)

func testmain() {

	fmt.Println("helo")
	address := ":8080"
	myHandler := http.NewServeMux()
	server := &http.Server{

		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Starting server on", address)
	log.Fatal(server.ListenAndServe())
}
