package model

import (
	"github.com/czlingo/wormhole/internal/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DSN_DEFAULT = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

var (
	db *gorm.DB
)

func New(config *define.Config) (*gorm.DB, error) {
	gormDB, err := newDB()
	if err != nil {
		return nil, err
	}
	db = gormDB

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}

func newDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      DSN_DEFAULT,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
	}))

	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
