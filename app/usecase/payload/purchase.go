package payload

import (
	"time"

	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

type Purchase struct {
	ID       valueobject.PurchaseID
	ItemID   valueobject.ItemID
	Quantity uint64
	BoughtAt time.Time
}

type PurchaseRequest struct {
	ItemID   valueobject.ItemID
	Quantity uint64
}
