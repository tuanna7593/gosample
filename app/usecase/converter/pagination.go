package converter

import (
	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func ConvertPaginationPayloadToValueObject(pl payload.PaginationRequest) valueobject.PaginationRequest {
	return valueobject.PaginationRequest{
		Page:  pl.Page,
		Limit: pl.Limit,
	}
}
