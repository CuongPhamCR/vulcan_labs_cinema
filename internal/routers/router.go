package routers

import (
	"fmt"
	"vulcan_labs_cinema/global"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	var r *gin.Engine

	fmt.Println("server mode :: ", global.Config.Server.Mode)
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// health-check
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})

	MainGroup := r.Group("/v1")

	RouterGroupApp.Cinema.InitCinemaRouter(MainGroup)

	return r
}
