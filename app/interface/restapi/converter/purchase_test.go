package converter

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/interface/restapi/presenter"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func TestConvertPurchasePayloadToResponse(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		pl := payload.Purchase{
			ID:       valueobject.PurchaseID(1),
			ItemID:   valueobject.ItemID(2),
			Quantity: 1,
			BoughtAt: time.Date(2021, 10, 16, 0, 0, 0, 0, time.Local),
		}
		got := ConvertPurchasePayloadToResponse(pl)
		want := presenter.Purchase{
			ID:       valueobject.PurchaseID(1),
			ItemID:   valueobject.ItemID(2),
			Quantity: 1,
			BoughtAt: time.Date(2021, 10, 16, 0, 0, 0, 0, time.Local).Unix(),
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})
}
