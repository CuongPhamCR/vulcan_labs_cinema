package utils

import (
	"fmt"
	"vulcan_labs_cinema/internal/interfaces"
)

func ShowCinema(cinema *interfaces.Cinema) {
	fmt.Println(`========= CINEMA =========`)
	for i := 0; i < cinema.Rows; i++ {
		for j := 0; j < cinema.Cols; j++ {
			seat := cinema.Seats[i][j]
			if seat.Taken {
				fmt.Print(`[X]`)
			} else {
				fmt.Print(`[ ]`)
			}
		}
		fmt.Println()
	}
	fmt.Println(`==========================`)
}
