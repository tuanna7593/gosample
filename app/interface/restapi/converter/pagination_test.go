package converter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tuanna7593/gosample/app/interface/restapi/presenter"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func TestConvertPaginationRequestToPayload(t *testing.T) {
	t.Run("#1 Success", func(t *testing.T) {
		t.Parallel()
		req := presenter.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		got := ConvertPaginationRequestToPayload(req)
		want := payload.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})
}
