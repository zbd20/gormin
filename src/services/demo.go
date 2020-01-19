package services

import (
	"github.com/jinzhu/gorm"
	"github.com/zbd20/gormin/src/models"
)

type demoService struct {
	db *gorm.DB
}

func newDemoService(db *gorm.DB) *demoService {
	return &demoService{db}
}

func (ds demoService) Login(user models.Login) (*models.Login, error) {
	var u models.Login

	err := ds.db.Model(&u).Where("username = ? AND password = ?", user.Username, user.Password).Find(&u).Error
	if err != nil {
		return &u, err
	}
	if err == gorm.ErrRecordNotFound {
		return &u, err
	}

	return &u, nil
}
