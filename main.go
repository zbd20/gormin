package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	log "github.com/zbd20/go-utils/blog"
	"github.com/zbd20/gormin/src/config"
	"github.com/zbd20/gormin/src/router"
)

var cnf = flag.String("config", "config.prod.yaml", "config path")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /gin/api/v1
func main() {
	log.InitLogs()
	defer log.CloseLogs()

	flag.Parse()
	if err := config.InitConfig(*cnf); err != nil {
		log.Fatalf("Init config error: %v", err)
	}

	app := router.NewRouter()

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		log.Fatal(err)
	}

}

func init() {
	cpus := os.Getenv("CPUS")
	if count, err := strconv.Atoi(cpus); err == nil {
		runtime.GOMAXPROCS(count)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}
