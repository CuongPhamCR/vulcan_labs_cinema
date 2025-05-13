package services

import (
	"vulcan_labs_cinema/global"
	"vulcan_labs_cinema/internal/interfaces"
	"vulcan_labs_cinema/pkg/utils"
)

type ICinemaService interface {
	InitCinema(rows int, cols int, minDistance int) error
}

type cinemaService struct{}

func NewCinemaService() ICinemaService {
	return &cinemaService{}
}

// InitCinema implements ICinemaService.
func (c *cinemaService) InitCinema(rows int, cols int, minDistance int) error {
	seats := make([][]*interfaces.Seat, rows)
	for i := range rows {
		seats[i] = make([]*interfaces.Seat, cols)
		for j := range cols {
			seats[i][j] = &interfaces.Seat{Row: i, Col: j, Taken: false, Group: 0}
		}
	}

	// reset global.Cinema
	global.Cinema = nil

	global.Cinema = &interfaces.Cinema{
		Rows:        rows,
		Cols:        cols,
		MinDistance: minDistance,
		Seats:       seats,
		NextGroupID: 1,
	}

	utils.ShowCinema(global.Cinema)

	return nil
}
