package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zbd20/gormin/src/models"
)

func NewMySQLClient(dbConfig models.DB) (*gorm.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DbName)
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.LogMode(dbConfig.Log)
	db.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Hour)

	return db, nil
}
