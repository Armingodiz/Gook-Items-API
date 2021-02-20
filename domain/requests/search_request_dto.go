package requests

import (
	"github.com/ArminGodiz/Gook-Items-API/domain/items"
	"github.com/ArminGodiz/Gook-Items-API/utils/rest_errors"
)

func (r *SearchItemRequest) Search() ([]*items.Item, *rest_errors.RestErr) {
	return nil, nil
}
