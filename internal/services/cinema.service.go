package services

import (
	"fmt"
	"sort"
	"vulcan_labs_cinema/global"
	"vulcan_labs_cinema/internal/interfaces"
	"vulcan_labs_cinema/pkg/response"
	"vulcan_labs_cinema/pkg/utils"
)

type ICinemaService interface {
	InitCinema(rows int, cols int, minDistance int) error
	GetAvailableSeats(count int) ([][]*interfaces.Seat, int)
	ReserveSeats(data *interfaces.ReserveSeatsInput) (seats []interfaces.Seat, errCode int, err error)
	CancelSeats(data *interfaces.CancelSeatsInput) (errCode int, err error)
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

// ReserveSeats implements ICinemaService.
func (c *cinemaService) ReserveSeats(data *interfaces.ReserveSeatsInput) (seats []interfaces.Seat, errCode int, err error) {
	cinema := global.Cinema

	// check cinema is initialized
	if cinema == nil {
		return nil, response.ErrCodeCinemaNotFound, nil
	}

	// Get group id
	cinema.GlobalMutex.Lock()
	groupID := cinema.NextGroupID
	cinema.NextGroupID++
	cinema.GlobalMutex.Unlock()

	// Seat data for response
	seats = []interfaces.Seat{}

	seats = []interfaces.Seat{}

	rowSet := map[int]struct{}{}
	for _, s := range data.Seats {
		if s.Row >= cinema.Rows || s.Col >= cinema.Cols {
			return nil, response.ErrCodeSeatNotAvailableOrInvalid, nil
		}
		rowSet[s.Row] = struct{}{}
	}

	lockedRows := []int{}
	for r := range rowSet {
		lockedRows = append(lockedRows, r)
	}
	sort.Ints(lockedRows)
	for _, r := range lockedRows {
		cinema.RowLocks[r].Lock()
	}
	defer func() {
		for _, r := range lockedRows {
			cinema.RowLocks[r].Unlock()
		}
	}()

	for _, s := range data.Seats {
		seat := cinema.Seats[s.Row][s.Col]
		if seat.IsBooked || !utils.IsValidSeat(s.Row, s.Col, cinema) {
			return nil, response.ErrCodeSeatNotAvailableOrInvalid, nil
		}
	}

	// Reserve seat
	for _, s := range data.Seats {
		seat := cinema.Seats[s.Row][s.Col]
		seat.IsBooked = true
		seat.Group = groupID
		global.Logger.Info(fmt.Sprintf("Reserved seat row: %d, col: %d", s.Row, s.Col))
		seats = append(seats, interfaces.Seat{Row: s.Row, Col: s.Col, Group: groupID, IsBooked: true})
	}

	// Show current cinema
	utils.ShowCinema(cinema)

	return seats, response.ErrCodeSuccess, nil
}

// CancelSeats implements ICinemaService.
func (c *cinemaService) CancelSeats(data *interfaces.CancelSeatsInput) (errCode int, err error) {
	cinema := global.Cinema

	// check cinema is initialized
	if cinema == nil {
		return response.ErrCodeCinemaNotFound, nil
	}

	rowSet := map[int]struct{}{}
	for _, s := range data.Seats {
		if s.Row >= cinema.Rows || s.Col >= cinema.Cols {
			return response.ErrCodeSeatNotFound, nil
		}
		rowSet[s.Row] = struct{}{}
	}

	lockedRows := []int{}
	for r := range rowSet {
		lockedRows = append(lockedRows, r)
	}

	sort.Ints(lockedRows)
	for _, r := range lockedRows {
		cinema.RowLocks[r].Lock()
	}
	defer func() {
		for _, r := range lockedRows {
			cinema.RowLocks[r].Unlock()
		}
	}()

	// Cancel
	for _, s := range data.Seats {
		seat := cinema.Seats[s.Row][s.Col]
		if seat == nil {
			return response.ErrCodeSeatNotFound, nil
		}
		if !seat.IsBooked {
			return response.ErrCodeSeatIsNotBooked, nil
		}

		seat.IsBooked = false
		seat.Group = 0

		global.Logger.Info(fmt.Sprintf("Cancelled seat row: %d, col: %d", s.Row, s.Col))
	}

	// Show current cinema
	utils.ShowCinema(cinema)

	return response.ErrCodeSuccess, nil
}
