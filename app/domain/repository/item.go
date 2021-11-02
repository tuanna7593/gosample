package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

type ItemRepository interface {
	AssignTx(txm TransactionManager)
	Create(ctx context.Context, item *entity.Item) error
	Updates(ctx context.Context, item *entity.Item, values map[string]interface{}) error
	List(ctx context.Context, pagination valueobject.PaginationRequest) ([]entity.Item, error)
	GetByID(ctx context.Context, itemID valueobject.ItemID) (entity.Item, error)
}
