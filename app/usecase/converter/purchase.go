package converter

import (
	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func ConvertPurchaseEntityToPayload(ent entity.Purchase) payload.Purchase {
	return payload.Purchase{
		ID:       ent.ID,
		ItemID:   ent.ItemID,
		Quantity: ent.Quantity,
		BoughtAt: ent.CreatedAt,
	}
}
