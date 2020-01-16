// +build dev

package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zbd20/gormin/docs"
)

func init() {
	swaggerURL := ginSwagger.URL("http://127.0.0.1:8100/swagger/doc.json")
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL)
}
