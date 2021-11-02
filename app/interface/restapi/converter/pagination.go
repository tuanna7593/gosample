package converter

import (
	"github.com/tuanna7593/gosample/app/interface/restapi/presenter"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func ConvertPaginationRequestToPayload(p presenter.PaginationRequest) payload.PaginationRequest {
	return payload.PaginationRequest{
		Page:  p.Page,
		Limit: p.Limit,
	}
}
