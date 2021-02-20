package items

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArminGodiz/Gook-Items-API/clients/redis_client"
	"github.com/ArminGodiz/Gook-Items-API/utils/rest_errors"
)

const (
	row = "items"
)

func (item *Item) Save() *rest_errors.RestErr {
	json1, err := json.Marshal(item)
	if err != nil {
		return rest_errors.NewBadRequestError("invalid form")
	}
	err = redis_client.DB.HSet(context.Background(), row, item.Id, json1).Err()
	if err != nil {
		return rest_errors.NewInternalServerError("error while saving in db" + err.Error())
	}
	return nil
}

func (item *Item) Get() (*Item, *rest_errors.RestErr) {
	at, err := redis_client.DB.HGet(context.Background(), row, item.Id).Result()
	if err != nil {
		return nil, rest_errors.NewInternalServerError("item found !")
	}
	fmt.Println(at)
	var result Item
	json.Unmarshal([]byte(at), &result)
	return &result, nil
}
