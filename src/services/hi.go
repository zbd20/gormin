package services

import (
	"github.com/jinzhu/gorm"
	"github.com/zbd20/gormin/src/models"
)

type hiService struct {
	db *gorm.DB
}

func newHiService(db *gorm.DB) *hiService {
	return &hiService{db}
}

func (h *hiService) Get() (models.Hi, error) {
	return models.Hi{
		Msg: "Gin Web Service is OK",
	}, nil
}
