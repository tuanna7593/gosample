package mysql

import (
	"github.com/tuanna7593/gosample/app/domain/repository"
	"gorm.io/gorm"
)

type TransactionManagerImpl struct {
	db *gorm.DB
}

// NewTxDataSQL is the contructor function
func NewTransactionManagerImpl() repository.TransactionManager {
	return &TransactionManagerImpl{}
}

func (tx *TransactionManagerImpl) Begin() {
	db := GetDB().Session(&gorm.Session{SkipDefaultTransaction: true})
	tx.db = db.Begin()
}

func (tx *TransactionManagerImpl) Commit() error {
	return tx.db.Commit().Error
}

func (tx *TransactionManagerImpl) Rollback() {
	tx.db.Rollback()
}

func (tx *TransactionManagerImpl) GetTx() interface{} {
	return tx.db
}
