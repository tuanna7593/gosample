package usecase

import (
	"context"

	"github.com/tuanna7593/gosample/app/usecase/payload"
)

type ItemUseCase interface {
	Create(ctx context.Context, item payload.CreateItemRequest) (payload.Item, error)
	List(ctx context.Context, pagination payload.PaginationRequest) ([]payload.Item, error)
	BuyItem(ctx context.Context, req payload.PurchaseRequest) (payload.Purchase, error)
}
