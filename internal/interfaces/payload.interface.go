package interfaces

type InitCinemaInput struct {
	Rows        int `json:"rows" binding:"required,gte=1"`
	Cols        int `json:"cols" binding:"required,gte=1"`
	MinDistance int `json:"min_distance" binding:"required,gte=1"`
}
