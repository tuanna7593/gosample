package entity

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

type Item struct {
	ID                valueobject.ItemID
	CreatedAt         time.Time
	TotalStockValue   uint64
	CurrentStockValue uint64
	SellingPrice      decimal.Decimal
}
