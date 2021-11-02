package payload

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

type CreateItemRequest struct {
	TotalStockValue uint64
	SellingPrice    decimal.Decimal
}

type Item struct {
	ID                valueobject.ItemID
	TotalStockValue   uint64
	CurrentStockValue uint64
	SellingPrice      decimal.Decimal
	PlacedAt          time.Time
}

type Items []Item
