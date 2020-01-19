package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zbd20/gormin/src/services"
)

type BaseController struct {
	db *gorm.DB
	bs *services.BaseService
	rg *gin.RouterGroup
}

func NewBaseController(eng *gin.Engine, db *gorm.DB) *BaseController {
	rg := eng.Group("/gin/api/v1")

	bc := &BaseController{
		db: db,
		bs: services.NewBaseService(db),
		rg: rg,
	}

	newHiController(bc)
	newDemoController(bc)

	return bc
}
