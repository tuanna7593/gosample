package entity

import (
	"time"

	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

type Purchase struct {
	ID        valueobject.PurchaseID
	CreatedAt time.Time
	ItemID    valueobject.ItemID
	Quantity  uint64
}
