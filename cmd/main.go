package main

import (
	"strconv"
	"vulcan_labs_cinema/global"
	"vulcan_labs_cinema/internal/initialize"
)

func main() {
	// Start service
	r := initialize.Run()
	// Run
	port := global.Config.Server.Port | 8000
	r.Run(":" + strconv.Itoa(port))

}
