package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/repository"
)

// PurchaseRepositoryImpl purchase repository implementation
type PurchaseRepositoryImpl struct {
	db *gorm.DB
}

func NewPurchaseRepositoryImpl() repository.PurchaseRepository {
	return &PurchaseRepositoryImpl{
		db: GetDB(),
	}
}

func (r *PurchaseRepositoryImpl) AssignTx(txm repository.TransactionManager) {
	tx := txm.GetTx().(*gorm.DB)
	r.db = tx
}

func (r *PurchaseRepositoryImpl) Create(ctx context.Context, purchase *entity.Purchase) error {
	return r.db.Create(purchase).Error
}
