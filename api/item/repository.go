package item

import (
	"context"

	"github.com/prayogatriady/restaurant-management/model/item_model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	ItemList(ctx context.Context) (items []*item_model.ItemList, err error)
	// AddItem(contact *model.Item) (err error)
	// EditItem(contact *model.Item) (err error)
	// DeleteItem(contactId int) (err error)

	GenDummyCategories(ctx context.Context, categories []*item_model.Category) (err error)
	GenDummyItems(ctx context.Context, items []*item_model.Item) (err error)
}

type itemRepository struct {
	Db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{
		Db: db,
	}
}

func (r *itemRepository) ItemList(ctx context.Context) (items []*item_model.ItemList, err error) {

	query := `
	select mi.id, mi.name, mi.description, mi.price, mi.quantity, mc.name as category_name, mi.is_active 
	from m_item mi
	join m_category mc 
	on mi.category_id = mc.id ;
	`
	err = r.Db.WithContext(ctx).Raw(query).Scan(&items).Error
	return

}

func (r *itemRepository) GenDummyItems(ctx context.Context, items []*item_model.Item) (err error) {
	err = r.Db.WithContext(ctx).Table("m_item").Create(&items).Error
	return
}

func (r *itemRepository) GenDummyCategories(ctx context.Context, categories []*item_model.Category) (err error) {
	err = r.Db.WithContext(ctx).Table("m_category").Create(&categories).Error
	return
}
