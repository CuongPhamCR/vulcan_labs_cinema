package global

import (
	"vulcan_labs_cinema/internal/interfaces"
	"vulcan_labs_cinema/pkg/logger"
	"vulcan_labs_cinema/pkg/setting"
)

var (
	Config          setting.Config
	Logger          *logger.LoggerZap
	Cinema          *interfaces.Cinema
	CurrentCinemaID int = 0
)
