package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/zbd20/go-utils/blog"
	_ "github.com/zbd20/gormin/docs"
	"github.com/zbd20/gormin/src/apis"
	"github.com/zbd20/gormin/src/config"
	"github.com/zbd20/gormin/src/db"
	"github.com/zbd20/gormin/src/middleware"
	"github.com/zbd20/gormin/src/models"
)

type Router struct {
	eng *gin.Engine
	bc  *apis.BaseController
}

var swagHandler gin.HandlerFunc

func NewRouter() *Router {
	serverConfig := config.GetConfig()
	mysqlClient, err := db.NewMySQLClient(serverConfig.DB)
	if err != nil {
		log.Fatalf("new mysql client error: %v\n", err)
		return nil
	}

	// running mode(debug/test/release)
	gin.SetMode(serverConfig.Mode)
	// 强制日志颜色化
	//gin.ForceConsoleColor()

	eng := gin.Default()

	eng.Use(middleware.Page(), middleware.LoggerToFile())

	bc := apis.NewBaseController(eng, mysqlClient)

	r := &Router{
		eng: eng,
		bc:  bc,
	}

	// 注册gorm回调
	models.RegisterCallbacks(mysqlClient)
	// 自动创建表
	models.AutoCreateTable(mysqlClient)

	if swagHandler != nil {
		eng.GET("/swagger/*any", swagHandler)
	}

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
