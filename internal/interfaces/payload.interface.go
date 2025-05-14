package interfaces

type InitCinemaInput struct {
	Rows        int `json:"rows" binding:"required,gte=1"`
	Cols        int `json:"cols" binding:"required,gte=1"`
	MinDistance int `json:"min_distance" binding:"required,gte=1"`
}

type SeatInput struct {
	Row int `json:"row" binding:"required,gte=0"`
	Col int `json:"col" binding:"required,gte=0"`
}

type ReserveSeatsInput struct {
	Seats []*SeatInput `json:"seats" binding:"required"`
}

type CancelSeatsInput struct {
	Seats []*SeatInput `json:"seats" binding:"required"`
}
