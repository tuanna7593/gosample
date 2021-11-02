package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type TransactionManager interface {
	Begin()
	Commit() error
	Rollback()
	GetTx() interface{}
}
