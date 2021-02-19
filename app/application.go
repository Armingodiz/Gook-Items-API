package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	MapUrls()
	srv := &http.Server{
		Handler:           router,
		Addr:              "127.0.0.1:3333",
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
