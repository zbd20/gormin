package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func RegisterCallbacks(db *gorm.DB) {
	//db.Callback().Create().Before("gorm:create").Register("gorm:update_create_time", createCallback)
	//db.Callback().Update().Before("gorm:update").Register("gorm:update_update_time", updateCallback)
	db.Callback().Create().Replace("gorm:update_time_stamp", createCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateCallback)
}

func createCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now()
		if ct, ok := scope.FieldByName("CreateTime"); ok {
			if ct.IsBlank {
				ct.Set(now)
			}
		}

		if ut, ok := scope.FieldByName("UpdateTime"); ok {
			if ut.IsBlank {
				ut.Set(now)
			}
		}
	}
}

func updateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now()
		if _, ok := scope.Get("gorm:update_time"); !ok {
			scope.SetColumn("UpdateTime", now)
		}
	}
}

func AutoCreateTable(db *gorm.DB) {
	tables := []interface{}{User{}}
	for _, t := range tables {
		table := db.HasTable(t)
		if !table {
			db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(t)
		}
	}
}
