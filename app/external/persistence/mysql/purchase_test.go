package mysql

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/tuanna7593/gosample/app/domain/entity"
	"github.com/tuanna7593/gosample/app/domain/valueobject"
	"github.com/tuanna7593/gosample/app/utils/testsupport"
)

func TestPurchaseRepositoryImpl_Create(t *testing.T) {
	t.Run("#1: Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		purchase := entity.Purchase{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}

		insertQuery := regexp.QuoteMeta("INSERT INTO `purchases` (`created_at`,`item_id`,`quantity`) VALUES (?,?,?)")
		mock.ExpectBegin()
		mock.ExpectExec(insertQuery).WillReturnResult(
			sqlmock.NewResult(1, 1),
		)
		mock.ExpectCommit()

		repo := PurchaseRepositoryImpl{
			db: db,
		}

		err = repo.Create(context.Background(), &purchase)
		if err != nil {
			t.Errorf("repo.Create() return an error:%v - want:nil", err)
			return
		}

		if purchase.ID == 0 {
			t.Errorf("ID of a new Purchase must be different zero:%d", purchase.ID)
			return
		}

		if purchase.CreatedAt.IsZero() {
			t.Error("CreatedAt of a new Purchase must be different zero time")
			return
		}

		want := entity.Purchase{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}

		if diff := cmp.Diff(
			purchase, want,
			cmpopts.IgnoreFields(entity.Purchase{}, "ID", "CreatedAt")); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("#2: Faield to create", func(t *testing.T) {
		t.Parallel()
		db, mock, err := testsupport.OpenDBConnection()
		if err != nil {
			panic(err)
		}

		purchase := entity.Purchase{
			ItemID:   valueobject.ItemID(1),
			Quantity: 2,
		}

		wannaErr := errors.New("failed to create purchase")
		insertQuery := regexp.QuoteMeta("INSERT INTO `purchases` (`created_at`,`item_id`,`quantity`) VALUES (?,?,?)")
		mock.ExpectBegin()
		mock.ExpectExec(insertQuery).WillReturnError(wannaErr)
		mock.ExpectRollback()

		repo := PurchaseRepositoryImpl{
			db: db,
		}

		err = repo.Create(context.Background(), &purchase)
		if !errors.Is(err, wannaErr) {
			t.Errorf("repo.Create() return an error:%v - want:%v", err, wannaErr)
			return
		}
	})
}
