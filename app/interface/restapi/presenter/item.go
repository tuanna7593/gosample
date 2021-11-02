package presenter

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"

	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

// CreateItemRequest the presenter for create Items
type CreateItemRequest struct {
	TotalStockValue uint64          `json:"total_stock_value" validate:"min=1"`
	SellingPrice    decimal.Decimal `json:"selling_price" validate:"monetary"`
}

// Validate check the request is valid
func (p CreateItemRequest) Validate() error {
	v := validator.New()
	v.RegisterCustomTypeFunc(validateDecimalType, decimal.Decimal{})
	if err := v.RegisterValidation("monetary", validateMonetary); err != nil {
		return err
	}

	if err := v.Struct(p); err != nil {
		switch e := err.(type) {
		case validator.ValidationErrors:
			errs := make(payload.Errors, 0, len(e))
			for _, ee := range e {
				switch f := ee.Field(); {
				case f == "TotalStockValue":
					errs = append(errs, payload.Error{
						Code:    payload.ErrCodeInvalidTotalStockValue,
						Message: "'total_stock_value' should be greater than 0",
						Param:   p.TotalStockValue,
						Type:    payload.ErrorTypeInvalidArgument,
					})
				case f == "SellingPrice":
					errs = append(errs, payload.Error{
						Code:    payload.ErrCodeInvalidTotalStockValue,
						Message: "'selling_price' should be a positive decimal value to two decimal places",
						Param:   p.TotalStockValue,
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

func validateDecimalType(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(decimal.Decimal); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}

	return nil
}

func validateMonetary(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	sellingPrice, err := decimal.NewFromString(val)
	if err != nil {
		return false
	}

	if !ok || !sellingPrice.GreaterThan(decimal.Zero) {
		return false
	}

	sellingPriceStrs := strings.Split(sellingPrice.String(), ".")
	if len(sellingPriceStrs) > 1 {
		if len(sellingPriceStrs[1]) > 2 {
			return false
		}
	}

	return true
}

type ItemResponse struct {
	ID                valueobject.ItemID `json:"id"`
	PlacedAt          int64              `json:"placed_at"`
	TotalStockValue   uint64             `json:"total_stock_value"`
	CurrentStockValue uint64             `json:"current_stock_value"`
	SellingPrice      decimal.Decimal    `json:"selling_price"`
}

type BuyItemRequest struct {
	Quantity uint64 `json:"quantity" validate:"min=1"`
}

// Validate check the request is valid
func (p BuyItemRequest) Validate() error {
	if err := validator.New().Struct(p); err != nil {
		return payload.Error{
			Code:    payload.ErrCodeInvalidBuyQuantity,
			Message: "'quantity' should be greater than 0",
			Param:   p.Quantity,
			Type:    payload.ErrorTypeInvalidArgument,
		}
	}

	return nil
}
