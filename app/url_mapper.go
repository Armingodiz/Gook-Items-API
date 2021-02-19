package app

import (
	"github.com/ArminGodiz/Gook-Items-API/controllers"
	"net/http"
)

func MapUrls() {
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)

}
