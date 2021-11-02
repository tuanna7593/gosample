package converter

import (
	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

// ConvertCreateItemRequestToEntity convert create item request payload to item entity
func ConvertCreateItemRequestToEntity(request payload.CreateItemRequest) entity.Item {
	return entity.Item{
		TotalStockValue:   request.TotalStockValue,
		CurrentStockValue: request.TotalStockValue,
		SellingPrice:      request.SellingPrice,
	}
}

// ConvertItemEntityToPayload convert item entity to payload
func ConvertItemEntityToPayload(item entity.Item) payload.Item {
	return payload.Item{
		ID:                item.ID,
		PlacedAt:          item.CreatedAt,
		TotalStockValue:   item.TotalStockValue,
		CurrentStockValue: item.CurrentStockValue,
		SellingPrice:      item.SellingPrice,
	}
}
