package presenter

import (
	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

type Purchase struct {
	ID       valueobject.PurchaseID `json:"id"`
	ItemID   valueobject.ItemID     `json:"item_id"`
	Quantity uint64                 `json:"quantity"`
	BoughtAt int64                  `json:"bought_at"`
}
