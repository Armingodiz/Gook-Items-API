package requests

import (
	"github.com/ArminGodiz/Gook-Items-API/utils/rest_errors"
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

func (r *SearchItemRequest) Compare() bool{
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
