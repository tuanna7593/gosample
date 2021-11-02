package converter

import (
	"github.com/tuanna7593/gosample/app/interface/restapi/presenter"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func ConvertPurchasePayloadToResponse(pl payload.Purchase) presenter.Purchase {
	return presenter.Purchase{
		ID:       pl.ID,
		ItemID:   pl.ItemID,
		Quantity: pl.Quantity,
		BoughtAt: pl.BoughtAt.Unix(),
	}
}
