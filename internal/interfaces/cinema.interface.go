package interfaces

import "sync"

type Seat struct {
	Row   int  `json:"row"`
	Col   int  `json:"col"`
	Group int  `json:"group"`
	Taken bool `json:"taken"`
}

type Cinema struct {
	Rows        int
	Cols        int
	MinDistance int
	Seats       [][]*Seat
	NextGroupID int
	Mutex       sync.Mutex
}
