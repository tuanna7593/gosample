package mysql

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/utils/testsupport"
)

func TestItemRepositoryImpl_Create(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		item := entity.Item{
			TotalStockValue:   1,
			CurrentStockValue: 1,
			SellingPrice:      decimal.NewFromFloat32(1.5),
		}

		insertQuery := regexp.QuoteMeta("INSERT INTO `items` (`created_at`,`total_stock_value`,`current_stock_value`,`selling_price`) VALUES (?,?,?,?)")
		mock.ExpectBegin()
		mock.ExpectExec(insertQuery).WillReturnResult(
			sqlmock.NewResult(1, 1),
		)
		mock.ExpectCommit()

		repo := ItemRepositoryImpl{
			db: db,
		}

		err = repo.Create(context.Background(), &item)
		if err != nil {
			t.Errorf("repo.Create() return an error:%v - want:nil", err)
			return
		}

		if item.ID == 0 {
			t.Errorf("ID of a new Item must be different zero:%d", item.ID)
			return
		}

		if item.CreatedAt.IsZero() {
			t.Errorf("CreatedAt of a new Item must be different zero time:%v", item.CreatedAt.IsZero())
			return
		}

		want := entity.Item{
			TotalStockValue:   1,
			CurrentStockValue: 1,
			SellingPrice:      decimal.NewFromFloat32(1.5),
		}

		if diff := cmp.Diff(
			item, want,
			cmpopts.IgnoreFields(entity.Item{}, "ID", "CreatedAt")); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#2: Failed when create item ", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		item := entity.Item{
			TotalStockValue:   1,
			CurrentStockValue: 1,
			SellingPrice:      decimal.NewFromFloat32(1.5),
		}

		wannaErr := errors.New("cannot conntect db")

		insertQuery := regexp.QuoteMeta("INSERT INTO `items` (`created_at`,`total_stock_value`,`current_stock_value`,`selling_price`) VALUES (?,?,?,?)")
		mock.ExpectBegin()
		mock.ExpectExec(insertQuery).WithArgs().WillReturnError(wannaErr)
		mock.ExpectRollback()

		repo := ItemRepositoryImpl{
			db: db,
		}

		err = repo.Create(context.Background(), &item)
		if !errors.Is(err, wannaErr) {
			t.Errorf("repo.Create() return an error:%v - want:%v", err, wannaErr)
		}
	})
}

func TestItemRepositoryImpl_List(t *testing.T) {
	t.Run("#1: List without paginaiton", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}
		repo := ItemRepositoryImpl{
			db: db,
		}

		selectQuery := regexp.QuoteMeta("SELECT * FROM `items`")
		mock.ExpectQuery(selectQuery).WillReturnRows(
			sqlmock.NewRows([]string{
				"id", "created_at", "total_stock_value",
				"current_stock_value", "selling_price"}).
				AddRow(1, time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local), 5, 4, decimal.NewFromFloat(1.55)).
				AddRow(2, time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local), 1, 1, decimal.NewFromFloat(2.55)),
		)

		got, err := repo.List(context.Background(), valueobject.PaginationRequest{})
		if err != nil {
			t.Errorf("repo.List() return an error:%v - want: nil", err)
			return
		}

		want := []entity.Item{
			{
				ID:                1,
				CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
				TotalStockValue:   5,
				CurrentStockValue: 4,
				SellingPrice:      decimal.NewFromFloat(1.55),
			},
			{
				ID:                2,
				CreatedAt:         time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local),
				TotalStockValue:   1,
				CurrentStockValue: 1,
				SellingPrice:      decimal.NewFromFloat(2.55),
			},
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#2: List with paginaiton has Page less than 1", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}
		repo := ItemRepositoryImpl{
			db: db,
		}

		selectQuery := regexp.QuoteMeta("SELECT * FROM `items`")
		mock.ExpectQuery(selectQuery).WillReturnRows(
			sqlmock.NewRows([]string{
				"id", "created_at", "total_stock_value",
				"current_stock_value", "selling_price"}).
				AddRow(1, time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local), 5, 4, decimal.NewFromFloat(1.55)).
				AddRow(2, time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local), 1, 1, decimal.NewFromFloat(2.55)),
		)

		pagination := valueobject.PaginationRequest{
			Page:  -1,
			Limit: 1,
		}
		got, err := repo.List(context.Background(), pagination)
		if err != nil {
			t.Errorf("repo.List() return an error:%v - want: nil", err)
			return
		}

		want := []entity.Item{
			{
				ID:                1,
				CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
				TotalStockValue:   5,
				CurrentStockValue: 4,
				SellingPrice:      decimal.NewFromFloat(1.55),
			},
			{
				ID:                2,
				CreatedAt:         time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local),
				TotalStockValue:   1,
				CurrentStockValue: 1,
				SellingPrice:      decimal.NewFromFloat(2.55),
			},
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#3: List with paginaiton has Limit less than 1", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}
		repo := ItemRepositoryImpl{
			db: db,
		}

		selectQuery := regexp.QuoteMeta("SELECT * FROM `items`")
		mock.ExpectQuery(selectQuery).WillReturnRows(
			sqlmock.NewRows([]string{
				"id", "created_at", "total_stock_value",
				"current_stock_value", "selling_price"}).
				AddRow(1, time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local), 5, 4, decimal.NewFromFloat(1.55)).
				AddRow(2, time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local), 1, 1, decimal.NewFromFloat(2.55)),
		)

		pagination := valueobject.PaginationRequest{
			Page:  1,
			Limit: 0,
		}
		got, err := repo.List(context.Background(), pagination)
		if err != nil {
			t.Errorf("repo.List() return an error:%v - want: nil", err)
			return
		}

		want := []entity.Item{
			{
				ID:                1,
				CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
				TotalStockValue:   5,
				CurrentStockValue: 4,
				SellingPrice:      decimal.NewFromFloat(1.55),
			},
			{
				ID:                2,
				CreatedAt:         time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local),
				TotalStockValue:   1,
				CurrentStockValue: 1,
				SellingPrice:      decimal.NewFromFloat(2.55),
			},
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#4: List with valid paginaiton", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}
		repo := ItemRepositoryImpl{
			db: db,
		}

		selectQuery := regexp.QuoteMeta("SELECT * FROM `items` LIMIT 5 OFFSET 5")
		mock.ExpectQuery(selectQuery).WillReturnRows(
			sqlmock.NewRows([]string{
				"id", "created_at", "total_stock_value",
				"current_stock_value", "selling_price"}).
				AddRow(1, time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local), 5, 4, decimal.NewFromFloat(1.55)).
				AddRow(2, time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local), 1, 1, decimal.NewFromFloat(2.55)),
		)

		pagination := valueobject.PaginationRequest{
			Page:  2,
			Limit: 5,
		}
		got, err := repo.List(context.Background(), pagination)
		if err != nil {
			t.Errorf("repo.List() return an error:%v - want: nil", err)
		}

		want := []entity.Item{
			{
				ID:                1,
				CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
				TotalStockValue:   5,
				CurrentStockValue: 4,
				SellingPrice:      decimal.NewFromFloat(1.55),
			},
			{
				ID:                2,
				CreatedAt:         time.Date(2021, 10, 17, 10, 0, 0, 0, time.Local),
				TotalStockValue:   1,
				CurrentStockValue: 1,
				SellingPrice:      decimal.NewFromFloat(2.55),
			},
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#5: Failed when get items", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}
		repo := ItemRepositoryImpl{
			db: db,
		}

		selectQuery := regexp.QuoteMeta("SELECT * FROM `items` LIMIT 5 OFFSET 5")
		wannaErr := errors.New("failed to get items")
		mock.ExpectQuery(selectQuery).WillReturnError(wannaErr)

		pagination := valueobject.PaginationRequest{
			Page:  2,
			Limit: 5,
		}
		got, err := repo.List(context.Background(), pagination)
		if !errors.Is(err, wannaErr) {
			t.Errorf("repo.List() return an error:%v - want: %v", err, wannaErr)
			return
		}

		var want []entity.Item
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})
}

func TestItemRepositoryImpl_GetByID(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		query := regexp.QuoteMeta("SELECT * FROM `items` WHERE `items`.id = ?")
		mock.ExpectQuery(query).WillReturnRows(
			sqlmock.NewRows([]string{
				"id", "total_stock_value", "current_stock_value",
				"selling_price", "created_at",
			}).AddRow(1, 5, 4, decimal.NewFromFloat(1.55), time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local)),
		)

		repo := ItemRepositoryImpl{
			db: db,
		}
		got, err := repo.GetByID(context.Background(), valueobject.ItemID(1))
		if err != nil {
			t.Errorf("repo.GetByID() return an error:%v - want:nil", err)
			return
		}

		want := entity.Item{
			ID:                valueobject.ItemID(1),
			TotalStockValue:   5,
			CurrentStockValue: 4,
			SellingPrice:      decimal.NewFromFloat32(1.55),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#2: Not found item", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		query := regexp.QuoteMeta("SELECT * FROM `items` WHERE `items`.id = ?")
		mock.ExpectQuery(query).WillReturnError(gorm.ErrRecordNotFound)

		repo := ItemRepositoryImpl{
			db: db,
		}
		got, err := repo.GetByID(context.Background(), valueobject.ItemID(1))
		if err != nil {
			t.Errorf("repo.GetByID() return an error:%v - want:nil", err)
			return
		}

		var want entity.Item
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#3: Failed to get item", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		query := regexp.QuoteMeta("SELECT * FROM `items` WHERE `items`.id = ?")
		wannaErr := errors.New("failed to get item")
		mock.ExpectQuery(query).WillReturnError(wannaErr)

		repo := ItemRepositoryImpl{
			db: db,
		}
		got, err := repo.GetByID(context.Background(), valueobject.ItemID(1))
		if !errors.Is(err, wannaErr) {
			t.Errorf("repo.GetByID() return an error:%v - want:%v", err, wannaErr)
			return
		}

		var want entity.Item
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})
}

func TestItemRepositoryImpl_Updates(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		item := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat32(1.55),
		}
		updateValues := map[string]interface{}{
			"current_stock_value": 4,
			"selling_price":       decimal.NewFromFloat(2.55),
		}

		updateQuery := regexp.QuoteMeta("UPDATE `items` SET `current_stock_value`=?,`selling_price`=? WHERE `id` = ?")
		mock.ExpectBegin()
		mock.ExpectExec(updateQuery).WithArgs(4, decimal.NewFromFloat(2.55), uint64(1)).WillReturnResult(
			sqlmock.NewResult(1, 1),
		)
		mock.ExpectCommit()

		repo := ItemRepositoryImpl{
			db: db,
		}

		err = repo.Updates(context.Background(), &item, updateValues)
		if err != nil {
			t.Errorf("repo.Updates() return an error:%v - want:nil", err)
			return
		}

		want := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 4,
			SellingPrice:      decimal.NewFromFloat32(2.55),
		}

		if diff := cmp.Diff(item, want); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#2: Failed to update item", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		item := entity.Item{
			ID:                valueobject.ItemID(1),
			CreatedAt:         time.Date(2021, 10, 16, 10, 0, 0, 0, time.Local),
			TotalStockValue:   5,
			CurrentStockValue: 5,
			SellingPrice:      decimal.NewFromFloat32(1.55),
		}
		updateValues := map[string]interface{}{
			"current_stock_value": 4,
			"selling_price":       decimal.NewFromFloat(2.55),
		}

		wannaErr := errors.New("failed to update")
		updateQuery := regexp.QuoteMeta("UPDATE `items` SET `current_stock_value`=?,`selling_price`=? WHERE `id` = ?")
		mock.ExpectBegin()
		mock.ExpectExec(updateQuery).WithArgs(4, decimal.NewFromFloat(2.55), uint64(1)).
			WillReturnError(wannaErr)
		mock.ExpectRollback()

		repo := ItemRepositoryImpl{
			db: db,
		}

		err = repo.Updates(context.Background(), &item, updateValues)
		if !errors.Is(err, wannaErr) {
			t.Errorf("repo.Updates() return an error:%v - want:%v", err, wannaErr)
			return
		}
	})
}
