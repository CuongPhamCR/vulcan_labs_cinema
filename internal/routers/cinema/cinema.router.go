package cinema

import (
	"vulcan_labs_cinema/internal/wire"

	"github.com/gin-gonic/gin"
)

type CinemaRouter struct{}

func (cr *CinemaRouter) InitCinemaRouter(Router *gin.RouterGroup) {
	cinemaController, _ := wire.InitCinemaRouterHandler()

	cinemaGroup := Router.Group("/cinema", cinemaController.InitCinema)
	{
		cinemaGroup.POST("/init")
		// cinemaGroup.GET("seats/available", cinemaController.GetAvailableSeats)
		// cinemaGroup.POST("seats/reserve", cinemaController.ReserveSeats)
		// cinemaGroup.POST("seats/cancel", cinemaController.CancelSeats)
	}
}
