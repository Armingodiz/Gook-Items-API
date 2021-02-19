package controllers

import (
	"encoding/json"
	"github.com/ArminGodiz/Gook-Items-API/domain/items"
	"github.com/ArminGodiz/Gook-Items-API/services"
	"github.com/ArminGodiz/Gook-Items-API/services/oauth"
	"github.com/ArminGodiz/Gook-Items-API/utils/http_utils"
	"io/ioutil"
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
		http_utils.RespondJson(w, err.Code, err)
		return
	}
	var itemRequest items.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.RespondJson(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		http_utils.RespondJson(w, http.StatusBadRequest, err)
		return
	}
	itemRequest.Seller = oauth.GetCallerId(r)
	result, err2 := services.ItemService.Create(itemRequest)
	if err2 != nil {
		http_utils.RespondJson(w, err2.Code, err)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
