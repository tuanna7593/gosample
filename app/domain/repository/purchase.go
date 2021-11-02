package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"github.com/tuanna7593/gosample/app/domain/entity"
)

type PurchaseRepository interface {
	AssignTx(txm TransactionManager)
	Create(ctx context.Context, purchase *entity.Purchase) error
}
