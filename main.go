package main

import (
	"log"
	"net/http"
	"practice.com/blog_service/internal/router"
	"time"
)

func main() {

	routerHandler := router.NewRouter()

	server := http.Server{
		Addr:              ":8000",
		Handler:           routerHandler,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
