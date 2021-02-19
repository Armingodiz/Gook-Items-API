package controller

import (
	"fmt"
	"github.com/ArminGodiz/Gook-Items-API/domain/items"
	"github.com/ArminGodiz/Gook-Items-API/services"
	"github.com/ArminGodiz/Gook-Items-API/services/oauth"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// Todo : return err to caller
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemService.Create(item)
	if err != nil {
		// TODO : return err to caller
	}
	fmt.Println(result)
}

func Get(w http.ResponseWriter, r *http.Request) {

}
