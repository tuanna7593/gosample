package converter

import (
	"github.com/tuanna7593/gosample/app/interface/restapi/presenter"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func ConvertCreateItemRequestToPayload(p presenter.CreateItemRequest) payload.CreateItemRequest {
	return payload.CreateItemRequest{
		TotalStockValue: p.TotalStockValue,
		SellingPrice:    p.SellingPrice,
	}
}

func ConvertPayloadItemToResponse(pl payload.Item) presenter.ItemResponse {
	return presenter.ItemResponse{
		ID:                pl.ID,
		PlacedAt:          pl.PlacedAt.Unix(),
		TotalStockValue:   pl.TotalStockValue,
		CurrentStockValue: pl.CurrentStockValue,
		SellingPrice:      pl.SellingPrice,
	}
}
