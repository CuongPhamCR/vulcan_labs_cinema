package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

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

var cinema *Cinema

func InitCinema(c *gin.Context) {
	var input struct {
		Rows        int `json:"rows"`
		Cols        int `json:"cols"`
		MinDistance int `json:"min_distance"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	seats := make([][]*Seat, input.Rows)
	for i := 0; i < input.Rows; i++ {
		seats[i] = make([]*Seat, input.Cols)
		for j := 0; j < input.Cols; j++ {
			seats[i][j] = &Seat{Row: i, Col: j, Taken: false, Group: 0}
		}
	}

	cinema = &Cinema{
		Rows:        input.Rows,
		Cols:        input.Cols,
		MinDistance: input.MinDistance,
		Seats:       seats,
		NextGroupID: 1,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cinema initialized"})
	ShowCinema()
}

func Distance(a, b *Seat) int {
	return int(math.Abs(float64(a.Row-b.Row)) + math.Abs(float64(a.Col-b.Col)))
}

func isFarEnough(row, col int) bool {
	for i := 0; i < cinema.Rows; i++ {
		for j := 0; j < cinema.Cols; j++ {
			seat := cinema.Seats[i][j]
			if seat.Taken && Distance(&Seat{Row: row, Col: col}, seat) < cinema.MinDistance {
				return false
			}
		}
	}
	return true
}

func GetAvailableSeats(c *gin.Context) {
	// count := 0
	fmt.Println("c . query :: ", c.Query("count"))
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count"})
		return
	}

	result := [][]*Seat{}
	for i := 0; i < cinema.Rows; i++ {
		for j := 0; j <= cinema.Cols-count; j++ {
			ok := true
			temp := []*Seat{}
			for k := 0; k < count; k++ {
				seat := cinema.Seats[i][j+k]
				if seat.Taken || !isFarEnough(seat.Row, seat.Col) {
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

	c.JSON(http.StatusOK, result)
}

func ReserveSeats(c *gin.Context) {
	var seats []Seat
	if err := c.BindJSON(&seats); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	cinema.Mutex.Lock()
	defer cinema.Mutex.Unlock()

	for _, s := range seats {
		if s.Row >= cinema.Rows || s.Col >= cinema.Cols || cinema.Seats[s.Row][s.Col].Taken || !isFarEnough(s.Row, s.Col) {
			c.JSON(http.StatusConflict, gin.H{"error": "Seat not available or violates distance"})
			return
		}
	}

	groupID := cinema.NextGroupID
	cinema.NextGroupID++

	for _, s := range seats {
		seat := cinema.Seats[s.Row][s.Col]
		seat.Taken = true
		seat.Group = groupID
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seats reserved", "group_id": groupID})
	ShowCinema()
}

func CancelSeats(c *gin.Context) {
	var seats []Seat
	if err := c.BindJSON(&seats); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	cinema.Mutex.Lock()
	defer cinema.Mutex.Unlock()

	for _, s := range seats {
		if s.Row < cinema.Rows && s.Col < cinema.Cols {
			seat := cinema.Seats[s.Row][s.Col]
			seat.Taken = false
			seat.Group = 0
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seats cancelled"})
	ShowCinema()
}

func ShowCinema() {
	fmt.Println(`===== CINEMA SEATING =====`)
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

func main() {
	r := gin.Default()

	r.POST("/cinema/init", InitCinema)
	r.GET("/seats/available", GetAvailableSeats)
	r.POST("/seats/reserve", ReserveSeats)
	r.POST("/seats/cancel", CancelSeats)

	r.Run(":8080")
}
