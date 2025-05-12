//go:build wireinject

package wire

import (
	"vulcan_labs_cinema/internal/controller"
	"vulcan_labs_cinema/internal/services"

	"github.com/google/wire"
)

func InitCinemaRouterHandler() (*controller.CinemaController, error) {
	wire.Build(services.NewCinemaService, controller.NewCinemaController)

	return new(controller.CinemaController), nil
}
