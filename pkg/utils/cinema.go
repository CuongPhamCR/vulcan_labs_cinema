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
	for i := range cinema.Rows {
		for j := range cinema.Cols {
			seat := cinema.Seats[i][j]
			if seat.IsBooked && ManhattanDistance(&interfaces.Seat{Row: row, Col: col}, seat) < cinema.MinDistance {
				return false
			}
		}
	}
	return true
}
