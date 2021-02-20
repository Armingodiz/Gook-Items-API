package requests

import (
	"fmt"
	"github.com/ArminGodiz/Gook-Items-API/domain/items"
	"github.com/ArminGodiz/Gook-Items-API/utils/rest_errors"
	"strconv"
	"strings"
)

type SearchItemRequest struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

var ItemAvailableFieldsForSearch = []string{"id", "seller", "title", "description", "price", "status"}

func (r *SearchItemRequest) Validate() *rest_errors.RestErr {
	if r.Value == "" {
		return rest_errors.NewBadRequestError("invalid input for value of field !")
	}
	if !isAvailableAsField(r.Field) {
		return rest_errors.NewBadRequestError("invalid field for search !")
	}
	return nil
}

func (r *SearchItemRequest) Compare(item items.Item) bool {
	switch r.Field {
	case "id":
		return item.Id == r.Value
	case "seller":
		return strings.EqualFold(strconv.Itoa(int(item.Seller)), r.Value)
	case "title":
		return item.Title == r.Value
	case "price":
		price := fmt.Sprintf("%f", item.Price)
		return price == r.Value
	case "status":
		return r.Value == item.Status
	}
	return false
}

func isAvailableAsField(field string) bool {
	for _, f := range ItemAvailableFieldsForSearch {
		if strings.EqualFold(f, field) {
			return true
		}
	}
	return false
}
