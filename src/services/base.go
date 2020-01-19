package services

import "github.com/jinzhu/gorm"

type BaseService struct {
	HiService   *hiService
	DemoService *demoService
}

func NewBaseService(db *gorm.DB) *BaseService {
	return &BaseService{
		HiService:   newHiService(db),
		DemoService: newDemoService(db),
	}
}
