package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zbd20/gormin/src/pkg"
)

type hiController struct {
	*BaseController
}

func newHiController(bc *BaseController) *hiController {
	hc := &hiController{bc}

	hc.rg.GET("hi", hc.Get)

	return hc
}

// @Summary Health check.
// @Description Health check.
// @Tags Hi
// @version 1.0
// @Accept json
// @Produce json
// @Success 200 {object} models.Hi OK
// @Failure 400 {string} string ERROR
// @Router /gin/api/v1/hi [get]
func (hc *hiController) Get(c *gin.Context) {
	result, err := hc.bs.HiService.Get()
	if err != nil {
		pkg.WriteResponse(c, pkg.HiErr, err)
		return
	}

	c.JSON(http.StatusOK, pkg.NewResult(result))
}
