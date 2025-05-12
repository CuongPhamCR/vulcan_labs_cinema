package initialize

import (
	"vulcan_labs_cinema/global"
	"vulcan_labs_cinema/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Log)
}
