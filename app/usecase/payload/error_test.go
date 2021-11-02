package payload

import (
	"testing"
)

func TestError_Error(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		err := Error{
			Code:    ErrCodeInvalidSellingPrice,
			Message: "error message test",
			Param:   nil,
			Type:    ErrorTypeBadRequest,
		}
		msg := err.Error()
		wannaMessage := "error message test"
		if msg != wannaMessage {
			t.Errorf("err.Error() return a message %s - want:%s", msg, wannaMessage)
		}
	})
}

func TestErrors_Error(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		errs := Errors{
			{
				Code:    ErrCodeInvalidSellingPrice,
				Message: "error message test1",
				Param:   nil,
				Type:    ErrorTypeBadRequest,
			},
			{
				Code:    ErrCodeInvalidTotalStockValue,
				Message: "error message test2",
				Param:   nil,
				Type:    ErrorTypeBadRequest,
			},
		}
		msg := errs.Error()
		wannaMessage := "error message test1\nerror message test2"
		if msg != wannaMessage {
			t.Errorf("err.Error() return a message %s - want:%s", msg, wannaMessage)
		}
	})
}
