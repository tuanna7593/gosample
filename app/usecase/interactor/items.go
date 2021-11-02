package interactor

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/repository"
	"github.com/tuanna7593/gosample/app/usecase"
	"github.com/tuanna7593/gosample/app/usecase/converter"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

// ItemUseCaseImpl implementation of Item usecase
type ItemUseCaseImpl struct {
	itemRepository     repository.ItemRepository
	purchaseRepository repository.PurchaseRepository
	txManager          repository.TransactionManager
}

// NewItemUseCaseInteractor create new instance of Item interactor
func NewItemUseCaseInteractor(
	itemRepo repository.ItemRepository,
	purchaseRepository repository.PurchaseRepository,
	txManager repository.TransactionManager,
) usecase.ItemUseCase {
	return &ItemUseCaseImpl{
		itemRepository:     itemRepo,
		purchaseRepository: purchaseRepository,
		txManager:          txManager,
	}
}

// Create create a new item
func (uc ItemUseCaseImpl) Create(ctx context.Context, request payload.CreateItemRequest) (payload.Item, error) {
	item := converter.ConvertCreateItemRequestToEntity(request)

	err := uc.itemRepository.Create(ctx, &item)
	if err != nil {
		return payload.Item{}, err
	}

	return converter.ConvertItemEntityToPayload(item), nil
}

// List get list item
func (uc ItemUseCaseImpl) List(ctx context.Context, pagination payload.PaginationRequest) ([]payload.Item, error) {
	paginationValueObject := converter.ConvertPaginationPayloadToValueObject(pagination)
	items, err := uc.itemRepository.List(ctx, paginationValueObject)
	if err != nil {
		log.Printf("failed to get items - pagination:%+v", paginationValueObject)
		return nil, err
	}

	itemResps := make([]payload.Item, len(items))
	for i := range items {
		itemResps[i] = converter.ConvertItemEntityToPayload(items[i])
	}

	return itemResps, nil
}

// List get list item
func (uc ItemUseCaseImpl) BuyItem(ctx context.Context, req payload.PurchaseRequest) (payload.Purchase, error) {
	// start transaction
	uc.txManager.Begin()

	// assign tx to repositories
	uc.itemRepository.AssignTx(uc.txManager)
	uc.purchaseRepository.AssignTx(uc.txManager)

	var err error
	defer func() {
		if err != nil {
			log.Printf("found error - rollback transaction:%v\n", err)
			uc.txManager.Rollback()
		}
	}()

	// find item
	item, err := uc.itemRepository.GetByID(ctx, req.ItemID)
	if err != nil {
		log.Printf("failed to get item:%d\n", req.ItemID)
		return payload.Purchase{}, err
	}

	if reflect.DeepEqual(item, entity.Item{}) {
		// not found item
		msg := fmt.Sprintf("not found item:%d", req.ItemID)
		log.Println(msg)
		err = payload.Error{
			Code:    payload.ErrCodeNotFoundItem,
			Message: msg,
			Param:   req.ItemID,
			Type:    payload.ErrorTypeBadRequest,
		}
		return payload.Purchase{}, err
	}

	// check the current stock value
	if item.CurrentStockValue < req.Quantity {
		msg := fmt.Sprintf(
			"the item out of stock - current quantity:%d - request quantity:%d",
			item.CurrentStockValue, req.Quantity,
		)
		log.Println(msg)
		err = payload.Error{
			Code:    payload.ErrCodeOutOfStock,
			Message: msg,
			Param:   req.Quantity,
			Type:    payload.ErrorTypeBadRequest,
		}
		return payload.Purchase{}, err
	}

	// update the stock value of item
	updateValues := map[string]interface{}{
		"current_stock_value": item.CurrentStockValue - req.Quantity,
	}
	err = uc.itemRepository.Updates(ctx, &item, updateValues)
	if err != nil {
		log.Printf("failed to update current stock of item:%d\n", item.ID)
		return payload.Purchase{}, err
	}

	// create purchase record
	purchaseEnt := entity.Purchase{
		ItemID:   req.ItemID,
		Quantity: req.Quantity,
	}
	err = uc.purchaseRepository.Create(ctx, &purchaseEnt)
	if err != nil {
		log.Printf("failed to create purchase:%+v\n", purchaseEnt)
		return payload.Purchase{}, err
	}

	// commit transaction
	errCommit := uc.txManager.Commit()
	if errCommit != nil {
		log.Printf("failed to commit transaction:%+v\n", errCommit)
		return payload.Purchase{}, errCommit
	}

	return converter.ConvertPurchaseEntityToPayload(purchaseEnt), nil
}
