package converter

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"

	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/interface/restapi/presenter"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func TestConvertCreateItemRequestToPayload(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		presenterReq := presenter.CreateItemRequest{
			TotalStockValue: 5,
			SellingPrice:    decimal.NewFromFloat(1.55),
		}
		payloadReq := ConvertCreateItemRequestToPayload(presenterReq)
		want := payload.CreateItemRequest{
			TotalStockValue: 5,
			SellingPrice:    decimal.NewFromFloat(1.55),
		}

		if diff := cmp.Diff(payloadReq, want); diff != "" {
			t.Error(diff)
		}
	})
}

func TestConvertPayloadItemToResponse(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		placedAt := time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local)
		pl := payload.Item{
			ID:                valueobject.ItemID(1),
			TotalStockValue:   5,
			CurrentStockValue: 4,
			SellingPrice:      decimal.NewFromFloat(1.55),
			PlacedAt:          placedAt,
		}
		resp := ConvertPayloadItemToResponse(pl)
		want := presenter.ItemResponse{
			ID:                valueobject.ItemID(1),
			PlacedAt:          placedAt.Unix(),
			TotalStockValue:   5,
			CurrentStockValue: 4,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}

		if diff := cmp.Diff(resp, want); diff != "" {
			t.Error(diff)
		}
	})
}
