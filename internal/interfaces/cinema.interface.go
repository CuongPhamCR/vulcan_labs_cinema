package interfaces

import "sync"

type Seat struct {
	Row      int  `json:"row"`
	Col      int  `json:"col"`
	Group    int  `json:"group"`
	IsBooked bool `json:"is_booked"`
}

type Cinema struct {
	Rows        int
	Cols        int
	MinDistance int
	Seats       [][]*Seat
	RowLocks    []sync.RWMutex
	NextGroupID int
	GlobalMutex sync.Mutex
}
