package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/external/persistence/mysql"
	"github.com/tuanna7593/gosample/app/interface/restapi/converter"
	"github.com/tuanna7593/gosample/app/interface/restapi/presenter"
	"github.com/tuanna7593/gosample/app/usecase/interactor"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

type ItemHandler struct {
	BaseHandler
}

// NewItemHandler create a new handler for Items
func NewItemHandler() *ItemHandler {
	return &ItemHandler{}
}

// Create create a new item
func (hdl *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req presenter.CreateItemRequest
		err error
	)

	defer func() {
		hdl.SetError(w, err)
	}()

	// decoding request body to struct
	if errDecode := json.NewDecoder(r.Body).Decode(&req); errDecode != nil {
		log.Printf("failed to decode request create item:%s\n", errDecode.Error())
		err = payload.Error{
			Message: "failed to decode create item request",
			Type:    payload.ErrorTypeBadRequest,
		}
		return
	}

	// validate user request
	err = req.Validate()
	if err != nil {
		log.Println("invalid create item request")
		return
	}

	// init payload request
	payloadRequest := converter.ConvertCreateItemRequestToPayload(req)

	// init usecase
	uc := interactor.NewItemUseCaseInteractor(
		mysql.NewItemRepositoryImpl(),
		nil,
		nil,
	)

	// execute use case create
	itemPayload, err := uc.Create(r.Context(), payloadRequest)
	if err != nil {
		return
	}

	// success
	resp := converter.ConvertPayloadItemToResponse(itemPayload)
	hdl.WriteResponse(w, http.StatusCreated, resp)
}

func (hdl *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	var (
		paginationRequest presenter.PaginationRequest
		err               error
	)

	defer func() {
		hdl.SetError(w, err)
	}()

	// parse pagination request
	err = paginationRequest.Parse(r.URL.Query())
	if err != nil {
		log.Println("failed to parse query string to pagination")
		return
	}

	// validate pagination request
	err = paginationRequest.Valiate()
	if err != nil {
		log.Printf("invalid pagination request:%+v\n", paginationRequest)
		return
	}

	// convert to payload
	payloadPagination := converter.ConvertPaginationRequestToPayload(paginationRequest)

	// init usecase
	uc := interactor.NewItemUseCaseInteractor(
		mysql.NewItemRepositoryImpl(),
		nil,
		nil,
	)

	items, err := uc.List(r.Context(), payloadPagination)
	if err != nil {
		log.Println("failed to get items")
		return
	}

	// convert payload to prenseter
	itemResp := make([]presenter.ItemResponse, len(items))
	for i := range itemResp {
		itemResp[i] = converter.ConvertPayloadItemToResponse(items[i])
	}

	// success
	hdl.WriteResponse(w, http.StatusOK, itemResp)
}

func (hdl *ItemHandler) BuyItem(w http.ResponseWriter, r *http.Request) {
	var (
		req presenter.BuyItemRequest
		err error
	)

	defer func() {
		hdl.SetError(w, err)
	}()

	itemIDStr := chi.URLParam(r, "item_id")
	if itemIDStr == "" {
		err = payload.Error{
			Code:    payload.ErrCodeInvalidItemID,
			Message: "not found item_id",
			Param:   nil,
			Type:    payload.ErrorTypeBadRequest,
		}
		return
	}

	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		err = payload.Error{
			Code:    payload.ErrCodeInvalidItemID,
			Message: "failed to parse item_id",
			Param:   itemIDStr,
			Type:    payload.ErrorTypeInvalidArgument,
		}
		return
	}

	// decoding request body to struct
	if errDecode := json.NewDecoder(r.Body).Decode(&req); errDecode != nil {
		log.Printf("failed to decode request buy item:%s\n", errDecode.Error())
		err = payload.Error{
			Message: "failed to decode buy item request",
			Type:    payload.ErrorTypeBadRequest,
		}
		return
	}

	// validate buy item request
	err = req.Validate()
	if err != nil {
		return
	}

	// init usecase
	uc := interactor.NewItemUseCaseInteractor(
		mysql.NewItemRepositoryImpl(),
		mysql.NewPurchaseRepositoryImpl(),
		mysql.NewTransactionManagerImpl(),
	)

	// execute use case
	purchase, err := uc.BuyItem(r.Context(), payload.PurchaseRequest{
		ItemID:   valueobject.ItemID(itemID),
		Quantity: req.Quantity,
	})
	if err != nil {
		fmt.Printf("failed to buy item:%+v\n", purchase)
		return
	}

	// success
	resp := converter.ConvertPurchasePayloadToResponse(purchase)
	hdl.WriteResponse(w, http.StatusCreated, resp)
}
