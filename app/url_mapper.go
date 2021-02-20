package app

import (
	"github.com/ArminGodiz/Gook-Items-API/controllers"
	"net/http"
)

func MapUrls() {
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemController.Get).Methods(http.MethodGet)
	router.HandleFunc("/ping", controllers.PingController.Pong).Methods(http.MethodGet)
}
