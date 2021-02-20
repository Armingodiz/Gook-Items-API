package requests

import (
	"context"
	"encoding/json"
	"github.com/ArminGodiz/Gook-Items-API/clients/redis_client"
	"github.com/ArminGodiz/Gook-Items-API/domain/items"
	"github.com/ArminGodiz/Gook-Items-API/utils/rest_errors"
)

func (r *SearchItemRequest) Search() ([]*items.Item, *rest_errors.RestErr) {
	allItems, err := redis_client.DB.HGetAll(context.Background(), "items").Result()
	if err != nil {
		return nil, rest_errors.NewInternalServerError("no item found !")
	}
	results := make([]*items.Item, 0)
	for _, temp := range allItems {
		var item items.Item
		err = json.Unmarshal([]byte(temp), &item)
		if err == nil {
			if r.Compare() {
				results = append(results, &item)
			}
		}
	}
	return results, nil
}
