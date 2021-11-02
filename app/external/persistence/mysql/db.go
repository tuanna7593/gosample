package mysql

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/tuanna7593/gosample/app/config"
)

var (
	once        sync.Once
	dbSingleton *gorm.DB
)

// initDB create a new DB instance
func InitDB(cfg config.MySQL) error {
	connectionStr := cfg.Conn()
	db, err := openDBConnection(connectionStr, gorm.Config{})
	if err != nil {
		err = fmt.Errorf("failed to connect database %s: %w", connectionStr, err)
		return err
	}

	gormer, err := db.DB()
	if err != nil {
		err = fmt.Errorf("failed to get database: %w", err)
		return err
	}

	gormer.SetMaxOpenConns(cfg.MaxOpenConns)
	gormer.SetMaxIdleConns(cfg.MaxIdleConns)

	return nil
}

// openDBConnection opens a DB connection.
func openDBConnection(connStr string, config gorm.Config) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, errConn := gorm.Open(mysql.Open(connStr), &config)
		if errConn != nil {
			err = errConn
			return
		}
		dbSingleton = db
	})
	if err != nil {
		return nil, err
	}
	return dbSingleton, nil
}

// GetDB gets the instance of singleton
func GetDB() *gorm.DB {
	return dbSingleton
}
