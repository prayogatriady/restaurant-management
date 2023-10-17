package item

import (
	"github.com/prayogatriady/restaurant-management/model"
)

type ItemRepository interface {
	GetItemList(pagination *model.Pagination) (contacts []*model.Item, totalRows int64, err error)
	AddItem(contact *model.Item) (err error)
	EditItem(contact *model.Item) (err error)
	DeleteItem(contactId int) (err error)
}

type itemRepository struct{}
