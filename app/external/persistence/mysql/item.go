package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/repository"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

// ItemRepositoryImpl item repository implementation
type ItemRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepositoryImpl() repository.ItemRepository {
	return &ItemRepositoryImpl{
		db: GetDB(),
	}
}

func (r *ItemRepositoryImpl) AssignTx(txm repository.TransactionManager) {
	tx := txm.GetTx().(*gorm.DB)
	r.db = tx
}

func (r *ItemRepositoryImpl) Create(ctx context.Context, item *entity.Item) error {
	return r.db.Create(item).Error
}

func (r *ItemRepositoryImpl) Updates(ctx context.Context, item *entity.Item, values map[string]interface{}) error {
	return r.db.Model(item).Updates(values).Error
}

func (r *ItemRepositoryImpl) List(ctx context.Context, pagination valueobject.PaginationRequest) ([]entity.Item, error) {
	var items []entity.Item
	err := r.db.Scopes(Paginate(pagination)).Find(&items).Error
	return items, err
}

func (r *ItemRepositoryImpl) GetByID(ctx context.Context, itemID valueobject.ItemID) (entity.Item, error) {
	var item entity.Item
	err := r.db.Take(&item, "`items`.id = ?", itemID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Item{}, nil
		}
		return entity.Item{}, err
	}

	return item, nil
}
