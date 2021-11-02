package mysql

import (
	"gorm.io/gorm"

	"github.com/tuanna7593/gosample/app/domain/valueobject"
)

// Paginate page with limit and offset
func Paginate(paging valueobject.PaginationRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paging.Page < 1 || paging.Limit < 1 {
			return db
		}

		offset := (paging.Page - 1) * paging.Limit
		return db.Limit(int(paging.Limit)).Offset(int(offset))
	}
}
