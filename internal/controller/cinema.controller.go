package controller

import (
	"net/http"
	"time"
	"vulcan_labs_cinema/internal/interfaces"
	"vulcan_labs_cinema/internal/services"
	"vulcan_labs_cinema/pkg/response"
	"vulcan_labs_cinema/pkg/utils/validation"

	"github.com/gin-gonic/gin"
)

type CinemaController struct {
	cinemaService services.ICinemaService
}

func NewCinemaController(
	cinemaService services.ICinemaService,
) *CinemaController {
	return &CinemaController{
		cinemaService: cinemaService,
	}
}

func (cc *CinemaController) InitCinema(c *gin.Context) {
	var payload interfaces.InitCinemaInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		errMessages := validation.FormatValidationError(err)
		response.ErrorResponse(c, http.StatusBadRequest, response.ErrCodeParamInvalid, errMessages[0])
		return
	}

	err := cc.cinemaService.InitCinema(payload.Rows, payload.Cols, payload.MinDistance)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, response.ErrCodeInternalServer, err.Error())
		return
	}

	type CinemaResponse struct {
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
		Rows        int    `json:"rows"`
		Cols        int    `json:"cols"`
		MinDistance int    `json:"min_distance"`
	}

	response.SuccessResponse(c, response.ErrCodeSuccess, CinemaResponse{
		Status:      "success",
		CreatedAt:   time.Now().Format(time.RFC3339),
		Rows:        payload.Rows,
		Cols:        payload.Cols,
		MinDistance: payload.MinDistance,
	})
}
