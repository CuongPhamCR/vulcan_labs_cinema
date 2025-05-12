package routers

import "vulcan_labs_cinema/internal/routers/cinema"

type RouterGroup struct {
	Cinema cinema.CinemaRouter
}

var RouterGroupApp = new(RouterGroup)
