package utils

import (
	"fmt"
	"math"
	"vulcan_labs_cinema/internal/interfaces"
)

func ShowCinema(cinema *interfaces.Cinema) {
	fmt.Println(`========= CINEMA =========`)
	for i := range cinema.Rows {
		for j := range cinema.Cols {
			seat := cinema.Seats[i][j]
			if seat.IsBooked {
				fmt.Print(`[X]`)
			} else {
				fmt.Print(`[ ]`)
			}
		}
		fmt.Println()
	}
	fmt.Println(`==========================`)
}

func ManhattanDistance(s1 *interfaces.Seat, s2 *interfaces.Seat) int {
	return int(math.Abs(float64(s1.Row-s2.Row)) + math.Abs(float64(s1.Col-s2.Col)))
}

func IsValidSeat(row, col int, cinema *interfaces.Cinema) bool {
	D := cinema.MinDistance
	for i := max(0, row-D); i <= min(cinema.Rows-1, row+D); i++ {
		for j := max(0, col-D); j <= min(cinema.Cols-1, col+D); j++ {
			if cinema.Seats[i][j].IsBooked && ManhattanDistance(&interfaces.Seat{Row: row, Col: col}, cinema.Seats[i][j]) < D {
				return false
			}
		}
	}
	return true
}
