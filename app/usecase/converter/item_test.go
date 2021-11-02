package converter

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func TestConvertCreateItemRequestToEntity(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		pl := payload.CreateItemRequest{
			TotalStockValue: 5,
			SellingPrice:    decimal.NewFromFloat(1.555),
		}
		itemEnt := ConvertCreateItemRequestToEntity(pl)
		want := entity.Item{
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.555),
		}

		if diff := cmp.Diff(itemEnt, want); diff != "" {
			t.Error(diff)
		}
	})
}

func TestConvertItemEntityToPayload(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		itemEnt := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   4,
			CurrentStockValue: 3,
			SellingPrice:      decimal.NewFromFloat(1.44),
		}

		itemPayload := ConvertItemEntityToPayload(itemEnt)
		want := payload.Item{
			ID:                valueobject.ItemID(1),
			PlacedAt:          time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   4,
			CurrentStockValue: 3,
			SellingPrice:      decimal.NewFromFloat(1.44),
		}

		if diff := cmp.Diff(itemPayload, want); diff != "" {
			t.Error(diff)
		}
	})
}
