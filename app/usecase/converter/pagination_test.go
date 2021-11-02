package converter

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func TestConvertPaginationPayloadToValueObject(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		pl := payload.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		got := ConvertPaginationPayloadToValueObject(pl)
		want := valueobject.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})
}
