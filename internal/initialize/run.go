package initialize

import (
	"vulcan_labs_cinema/global"
	"vulcan_labs_cinema/internal/routers"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// Load Config
	LoadConfig()

	// Init Logger
	InitLogger()

	global.Logger.Info("load config success to global.Config")

	// Init Router
	r := routers.NewRouter()

	return r
}
