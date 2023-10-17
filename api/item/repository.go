package item

import (
	"context"

	"github.com/prayogatriady/restaurant-management/model/item_model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	// ItemList(pagination *model.Pagination) (contacts []*model.Item, totalRows int64, err error)
	// AddItem(contact *model.Item) (err error)
	// EditItem(contact *model.Item) (err error)
	// DeleteItem(contactId int) (err error)

	AddBulkCategories(ctx context.Context, categories []*item_model.Category) (err error)
	AddBulkItems(ctx context.Context, items []*item_model.Item) (err error)
}

type itemRepository struct {
	Db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{
		Db: db,
	}
}

func (r *itemRepository) AddBulkItems(ctx context.Context, items []*item_model.Item) (err error) {
	err = r.Db.WithContext(ctx).Table("m_item").Create(&items).Error
	return
}

func (r *itemRepository) AddBulkCategories(ctx context.Context, categories []*item_model.Category) (err error) {
	err = r.Db.WithContext(ctx).Table("m_category").Create(&categories).Error
	return
}
