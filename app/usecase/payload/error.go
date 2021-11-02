package payload

import "strings"

type ErrorType string

const (
	ErrorTypeInvalidArgument ErrorType = "invalid argument"
	ErrorTypeNotFound        ErrorType = "not found"
	ErrorTypeBadRequest      ErrorType = "bad request"
)

type ErrorCode string

const (
	// error code of item
	ErrCodeInvalidItemID          ErrorCode = "ERR_INVALID_ITEM_ID"
	ErrCodeInvalidTotalStockValue ErrorCode = "ERR_INVALID_TOTAL_STOCK_VALUE"
	ErrCodeInvalidSellingPrice    ErrorCode = "ERR_INVALID_SELLING_PRICE"
	ErrCodeNotFoundItem           ErrorCode = "ERR_NOT_FOUMD_ITEM"

	// error code of pagination
	ErrCodeInvalidPage  ErrorCode = "ERR_INVALID_PAGE"
	ErrCodeInvalidLimit ErrorCode = "ERR_INVALID_LIMIT"

	// error code of buy item
	ErrCodeInvalidBuyQuantity ErrorCode = "ERR_INVALID_BUY_QUANTITY"
	ErrCodeOutOfStock         ErrorCode = "ERR_OUT_OF_STOCK"
)

type Error struct {
	Code    ErrorCode
	Message string
	Param   interface{}
	Type    ErrorType
}

func (e Error) Error() string {
	return e.Message
}

type Errors []Error

func (es Errors) Error() string {
	msg := []string{}
	for i := range es {
		msg = append(msg, es[i].Message)
	}

	return strings.Join(msg, "\n")
}
