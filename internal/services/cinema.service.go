package services

import (
	"fmt"
	"vulcan_labs_cinema/global"
	"vulcan_labs_cinema/internal/interfaces"
	"vulcan_labs_cinema/pkg/response"
	"vulcan_labs_cinema/pkg/utils"
)

type ICinemaService interface {
	InitCinema(rows int, cols int, minDistance int) error
	GetAvailableSeats(count int) ([][]*interfaces.Seat, int)
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
			seats[i][j] = &interfaces.Seat{Row: i, Col: j, IsBooked: false, Group: 0}
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

// GetAvailableSeats implements ICinemaService.
func (c *cinemaService) GetAvailableSeats(count int) (seats [][]*interfaces.Seat, errCode int) {
	fmt.Printf("count: %d,", count)
	result := [][]*interfaces.Seat{}

	cinema := global.Cinema

	// check cinema is initialized
	if cinema == nil {
		return nil, response.ErrCodeCinemaNotFound
	}

	for i := range cinema.Rows {
		for j := 0; j <= cinema.Cols-count; j++ {
			ok := true
			temp := []*interfaces.Seat{}
			for k := range count {
				seat := cinema.Seats[i][j+k]
				if seat.IsBooked || !utils.IsValidSeat(seat.Row, seat.Col, cinema) {
					ok = false
					break
				}
				temp = append(temp, seat)
			}
			if ok {
				result = append(result, temp)
			}
		}
	}

	return result, response.ErrCodeSuccess
}
