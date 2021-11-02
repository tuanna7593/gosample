package presenter

import (
	"log"
	"net/url"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/tuanna7593/gosample/app/usecase/payload"
)

type PaginationRequest struct {
	Page  int64 `validate:"min=1"`
	Limit int64 `validate:"min=1"`
}

func (p *PaginationRequest) Valiate() error {
	if err := validator.New().Struct(p); err != nil {
		switch e := err.(type) {
		case validator.ValidationErrors:
			errs := make(payload.Errors, 0, len(e))
			for _, ee := range e {
				switch f := ee.Field(); {
				case f == "Page":
					errs = append(errs, payload.Error{
						Code:    payload.ErrCodeInvalidPage,
						Message: "'page' should be greater than 0",
						Param:   p.Page,
						Type:    payload.ErrorTypeInvalidArgument,
					})
				case f == "Limit":
					errs = append(errs, payload.Error{
						Code:    payload.ErrCodeInvalidLimit,
						Message: "'limit' should be greater than 0",
						Param:   p.Limit,
						Type:    payload.ErrorTypeInvalidArgument,
					})
				}
			}
			return errs
		default:
			return err
		}
	}

	return nil
}

func (p *PaginationRequest) Parse(qs url.Values) error {
	// init default value
	p.Limit = 1
	p.Page = 1

	// parse from query string if exists
	if pageStr := qs.Get("page"); pageStr != "" {
		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			log.Printf("failed to parse page query to int64:%s\n", pageStr)
			return payload.Error{
				Code:    payload.ErrCodeInvalidPage,
				Message: "'page' should be an integer and greater than 0",
				Param:   pageStr,
				Type:    payload.ErrorTypeInvalidArgument,
			}
		}
		p.Page = page
	}

	if limitStr := qs.Get("limit"); limitStr != "" {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			log.Printf("failed to parse limit query to int64:%s\n", limitStr)
			return payload.Error{
				Code:    payload.ErrCodeInvalidLimit,
				Message: "'limit' should be an integer and greater than 0",
				Param:   limitStr,
				Type:    payload.ErrorTypeInvalidArgument,
			}
		}
		p.Limit = limit
	}

	return nil
}
