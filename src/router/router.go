package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	log "github.com/zbd20/go-utils/blog"
	_ "github.com/zbd20/gormin/docs"
	"github.com/zbd20/gormin/src/apis"
	"github.com/zbd20/gormin/src/config"
	"github.com/zbd20/gormin/src/db"
)

type Router struct {
	eng *gin.Engine
	bc  *apis.BaseController
}

func NewRouter() *Router {
	serverConfig := config.GetConfig()
	mysqlClient, err := db.NewMySQLClient(serverConfig.DB)
	if err != nil {
		log.Fatalf("new mysql client error: %v\n", err)
		return nil
	}

	// running mode(debug/test/release)
	gin.SetMode(serverConfig.Mode)
	eng := gin.Default()

	bc := apis.NewBaseController(eng, mysqlClient)

	r := &Router{
		eng: eng,
		bc:  bc,
	}

	swaggerURL := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	return r
}

func (r *Router) Run() error {
	serverConfig := config.GetConfig()
	log.Infof("start http server: %s", serverConfig.Addr)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", serverConfig.Addr),
		Handler:      r.eng,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}
