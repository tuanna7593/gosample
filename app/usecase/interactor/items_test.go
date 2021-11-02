package interactor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/repository/mock"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/usecase/payload"
)

func TestItemUseCaseImpl_Create(t *testing.T) {
	t.Run("#1 Success", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mItemRepo := mock.NewMockItemRepository(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository: mItemRepo,
		}
		ctx := context.Background()
		request := payload.CreateItemRequest{
			TotalStockValue: 5,
			SellingPrice:    decimal.NewFromFloat(1.55),
		}
		itemEntityRequest := entity.Item{
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}
		mItemRepo.EXPECT().Create(ctx, &itemEntityRequest).Return(nil)
		got, err := uc.Create(ctx, request)
		if err != nil {
			t.Errorf("uc.Create() return an error:%v - want:nil", err)
			return
		}

		want := payload.Item{
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}

		if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(payload.Item{}, "ID", "PlacedAt")); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#2 Failed when create item", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mItemRepo := mock.NewMockItemRepository(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository: mItemRepo,
		}
		ctx := context.Background()
		request := payload.CreateItemRequest{
			TotalStockValue: 5,
			SellingPrice:    decimal.NewFromFloat(1.55),
		}
		itemEntityRequest := entity.Item{
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}
		wannaErr := errors.New("failed to create a new item")
		mItemRepo.EXPECT().Create(ctx, &itemEntityRequest).Return(wannaErr)
		got, err := uc.Create(ctx, request)
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.Create() return an error:%v - want:%v", err, wannaErr)
			return
		}

		if diff := cmp.Diff(got, payload.Item{}); diff != "" {
			t.Error(diff)
		}
	})
}

func TestItemUseCaseImpl_List(t *testing.T) {
	t.Run("#1 Success", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mItemRepo := mock.NewMockItemRepository(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository: mItemRepo,
		}
		ctx := context.Background()
		paginationPl := payload.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		paginationVal := valueobject.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		itemEnts := []entity.Item{
			{
				ID:                valueobject.ItemID(1),
				CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
				TotalStockValue:   5,
				CurrentStockValue: 5,
				SellingPrice:      decimal.NewFromFloat(1.55),
			},
			{
				ID:                valueobject.ItemID(2),
				CreatedAt:         time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local),
				TotalStockValue:   10,
				CurrentStockValue: 5,
				SellingPrice:      decimal.NewFromFloat(3),
			},
		}
		mItemRepo.EXPECT().List(ctx, paginationVal).Return(itemEnts, nil)
		got, err := uc.List(ctx, paginationPl)
		if err != nil {
			t.Errorf("uc.List() return an error:%v - want:nil", err)
			return
		}

		want := []payload.Item{
			{
				ID:                valueobject.ItemID(1),
				PlacedAt:          time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
				TotalStockValue:   5,
				CurrentStockValue: 5,
				SellingPrice:      decimal.NewFromFloat(1.55),
			},
			{
				ID:                valueobject.ItemID(2),
				PlacedAt:          time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local),
				TotalStockValue:   10,
				CurrentStockValue: 5,
				SellingPrice:      decimal.NewFromFloat(3),
			},
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#2 Failed when get items", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mItemRepo := mock.NewMockItemRepository(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository: mItemRepo,
		}
		ctx := context.Background()
		paginationPl := payload.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		paginationVal := valueobject.PaginationRequest{
			Page:  1,
			Limit: 5,
		}
		wannaErr := errors.New("failed to get items")
		mItemRepo.EXPECT().List(ctx, paginationVal).Return(nil, wannaErr)
		got, err := uc.List(ctx, paginationPl)
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.List() return an error:%v - want:%v", err, wannaErr)
			return
		}

		var want []payload.Item
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})
}

func TestItemUseCaseImpl_BuyItem(t *testing.T) {
	t.Run("#1: Failed to get item", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mItemRepo := mock.NewMockItemRepository(mockCtrl)
		mPurchaseRepo := mock.NewMockPurchaseRepository(mockCtrl)
		mTxManager := mock.NewMockTransactionManager(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository:     mItemRepo,
			purchaseRepository: mPurchaseRepo,
			txManager:          mTxManager,
		}
		ctx := context.Background()
		req := payload.PurchaseRequest{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}
		wannaErr := errors.New("failed to get item")
		mTxManager.EXPECT().Begin()
		mItemRepo.EXPECT().AssignTx(mTxManager)
		mPurchaseRepo.EXPECT().AssignTx(mTxManager)
		mItemRepo.EXPECT().GetByID(ctx, req.ItemID).Return(entity.Item{}, wannaErr)
		mTxManager.EXPECT().Rollback()

		_, err := uc.BuyItem(ctx, req)
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.BuyItem() return an error:%v - want:%v", err, wannaErr)
		}
	})

	t.Run("#2: Not found item", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mItemRepo := mock.NewMockItemRepository(mockCtrl)
		mPurchaseRepo := mock.NewMockPurchaseRepository(mockCtrl)
		mTxManager := mock.NewMockTransactionManager(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository:     mItemRepo,
			purchaseRepository: mPurchaseRepo,
			txManager:          mTxManager,
		}
		ctx := context.Background()
		req := payload.PurchaseRequest{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}

		mTxManager.EXPECT().Begin()
		mItemRepo.EXPECT().AssignTx(mTxManager)
		mPurchaseRepo.EXPECT().AssignTx(mTxManager)
		mItemRepo.EXPECT().GetByID(ctx, req.ItemID).Return(entity.Item{}, nil)
		mTxManager.EXPECT().Rollback()

		_, err := uc.BuyItem(ctx, req)
		wannaErr := payload.Error{
			Code:    payload.ErrCodeNotFoundItem,
			Message: "not found item:1",
			Param:   req.ItemID,
			Type:    payload.ErrorTypeBadRequest,
		}
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.BuyItem() return an error:%v - want:%v", err, wannaErr)
		}
	})

	t.Run("#3: Item out of stock", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mItemRepo := mock.NewMockItemRepository(mockCtrl)
		mPurchaseRepo := mock.NewMockPurchaseRepository(mockCtrl)
		mTxManager := mock.NewMockTransactionManager(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository:     mItemRepo,
			purchaseRepository: mPurchaseRepo,
			txManager:          mTxManager,
		}
		ctx := context.Background()
		req := payload.PurchaseRequest{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}
		item := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 1,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}

		mTxManager.EXPECT().Begin()
		mItemRepo.EXPECT().AssignTx(mTxManager)
		mPurchaseRepo.EXPECT().AssignTx(mTxManager)
		mItemRepo.EXPECT().GetByID(ctx, req.ItemID).Return(item, nil)
		mTxManager.EXPECT().Rollback()

		_, err := uc.BuyItem(ctx, req)
		wannaErr := payload.Error{
			Code:    payload.ErrCodeOutOfStock,
			Message: "the item out of stock - current quantity:1 - request quantity:2",
			Param:   req.Quantity,
			Type:    payload.ErrorTypeBadRequest,
		}
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.BuyItem() return an error:%v - want:%v", err, wannaErr)
		}
	})

	t.Run("#4: Failed to update current stock of item", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mItemRepo := mock.NewMockItemRepository(mockCtrl)
		mPurchaseRepo := mock.NewMockPurchaseRepository(mockCtrl)
		mTxManager := mock.NewMockTransactionManager(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository:     mItemRepo,
			purchaseRepository: mPurchaseRepo,
			txManager:          mTxManager,
		}
		ctx := context.Background()
		req := payload.PurchaseRequest{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}
		item := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}
		updateValues := map[string]interface{}{
			"current_stock_value": uint64(3),
		}
		wannaErr := errors.New("failed to update item")

		mTxManager.EXPECT().Begin()
		mItemRepo.EXPECT().AssignTx(mTxManager)
		mPurchaseRepo.EXPECT().AssignTx(mTxManager)
		mItemRepo.EXPECT().GetByID(ctx, req.ItemID).Return(item, nil)
		mItemRepo.EXPECT().Updates(ctx, &item, updateValues).Return(wannaErr)
		mTxManager.EXPECT().Rollback()

		_, err := uc.BuyItem(ctx, req)
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.BuyItem() return an error:%v - want:%v", err, wannaErr)
		}
	})

	t.Run("#5: Failed to create purchase", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mItemRepo := mock.NewMockItemRepository(mockCtrl)
		mPurchaseRepo := mock.NewMockPurchaseRepository(mockCtrl)
		mTxManager := mock.NewMockTransactionManager(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository:     mItemRepo,
			purchaseRepository: mPurchaseRepo,
			txManager:          mTxManager,
		}
		ctx := context.Background()
		req := payload.PurchaseRequest{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}
		item := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}
		updateValues := map[string]interface{}{
			"current_stock_value": uint64(3),
		}
		purchaseEnt := entity.Purchase{
			ItemID:   req.ItemID,
			Quantity: req.Quantity,
		}
		wannaErr := errors.New("failed to update item")

		mTxManager.EXPECT().Begin()
		mItemRepo.EXPECT().AssignTx(mTxManager)
		mPurchaseRepo.EXPECT().AssignTx(mTxManager)
		mItemRepo.EXPECT().GetByID(ctx, req.ItemID).Return(item, nil)
		mItemRepo.EXPECT().Updates(ctx, &item, updateValues).Return(nil)
		mPurchaseRepo.EXPECT().Create(ctx, &purchaseEnt).Return(wannaErr)
		mTxManager.EXPECT().Rollback()

		_, err := uc.BuyItem(ctx, req)
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.BuyItem() return an error:%v - want:%v", err, wannaErr)
		}
	})

	t.Run("#6: Failed to commit transaction", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mItemRepo := mock.NewMockItemRepository(mockCtrl)
		mPurchaseRepo := mock.NewMockPurchaseRepository(mockCtrl)
		mTxManager := mock.NewMockTransactionManager(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository:     mItemRepo,
			purchaseRepository: mPurchaseRepo,
			txManager:          mTxManager,
		}
		ctx := context.Background()
		req := payload.PurchaseRequest{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}
		item := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}
		updateValues := map[string]interface{}{
			"current_stock_value": uint64(3),
		}
		purchaseEnt := entity.Purchase{
			ItemID:   req.ItemID,
			Quantity: req.Quantity,
		}
		wannaErr := errors.New("failed to commit transaction")

		mTxManager.EXPECT().Begin()
		mItemRepo.EXPECT().AssignTx(mTxManager)
		mPurchaseRepo.EXPECT().AssignTx(mTxManager)
		mItemRepo.EXPECT().GetByID(ctx, req.ItemID).Return(item, nil)
		mItemRepo.EXPECT().Updates(ctx, &item, updateValues).Return(nil)
		mPurchaseRepo.EXPECT().Create(ctx, &purchaseEnt).Return(nil)
		mTxManager.EXPECT().Commit().Return(wannaErr)

		_, err := uc.BuyItem(ctx, req)
		if !errors.Is(err, wannaErr) {
			t.Errorf("uc.BuyItem() return an error:%v - want:%v", err, wannaErr)
		}
	})

	t.Run("#7: Success", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mItemRepo := mock.NewMockItemRepository(mockCtrl)
		mPurchaseRepo := mock.NewMockPurchaseRepository(mockCtrl)
		mTxManager := mock.NewMockTransactionManager(mockCtrl)

		uc := ItemUseCaseImpl{
			itemRepository:     mItemRepo,
			purchaseRepository: mPurchaseRepo,
			txManager:          mTxManager,
		}
		ctx := context.Background()
		req := payload.PurchaseRequest{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}
		item := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat(1.55),
		}
		updateValues := map[string]interface{}{
			"current_stock_value": uint64(3),
		}
		purchaseEnt := entity.Purchase{
			ItemID:   req.ItemID,
			Quantity: req.Quantity,
		}

		mTxManager.EXPECT().Begin()
		mItemRepo.EXPECT().AssignTx(mTxManager)
		mPurchaseRepo.EXPECT().AssignTx(mTxManager)
		mItemRepo.EXPECT().GetByID(ctx, req.ItemID).Return(item, nil)
		mItemRepo.EXPECT().Updates(ctx, &item, updateValues).Return(nil)
		mPurchaseRepo.EXPECT().Create(ctx, &purchaseEnt).Return(nil)
		mTxManager.EXPECT().Commit().Return(nil)

		got, err := uc.BuyItem(ctx, req)
		if err != nil {
			t.Errorf("uc.BuyItem() return an error:%v - want:nil", err)
			return
		}
		wannaPurchase := payload.Purchase{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}
		if diff := cmp.Diff(
			got, wannaPurchase,
			cmpopts.IgnoreFields(payload.Purchase{}, "ID", "BoughtAt"),
		); diff != "" {
			t.Error(diff)
		}
	})
}
