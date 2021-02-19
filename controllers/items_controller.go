package controllers

import (
	"fmt"
	"github.com/ArminGodiz/Gook-Items-API/domain/items"
	"github.com/ArminGodiz/Gook-Items-API/services"
	"github.com/ArminGodiz/Gook-Items-API/services/oauth"
	"net/http"
)

var (
	ItemController itemControllerInterface = &itemController{}
)

type itemControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}
type itemController struct {
}

func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
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

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
