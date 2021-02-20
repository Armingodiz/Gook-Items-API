package services

import (
	"github.com/ArminGodiz/Gook-Items-API/domain/items"
	"github.com/ArminGodiz/Gook-Items-API/utils/rest_errors"
)

var (
	ItemService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}
type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	item := items.Item{Id: id}
	result, err := item.Get()
	if err != nil {
		return nil, err
	}
	return result, nil
}
