package apis

import (
	"github.com/zbd20/gormin/src/middleware"
)

type demoController struct {
	*BaseController
}

func newDemoController(bc *BaseController) *demoController {
	dc := &demoController{bc}

	amw := middleware.Jwt(bc.db)

	dc.rg.POST("/login", amw.LoginHandler)
	return dc
}
