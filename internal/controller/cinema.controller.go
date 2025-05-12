package controller

import (
	"net/http"
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

	cinemaId, err := cc.cinemaService.InitCinema(payload.Rows, payload.Cols, payload.MinDistance)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, response.ErrCodeInternalServer, err.Error())
		return
	}

	response.SuccessResponse(c, response.ErrCodeSuccess, map[string]interface{}{
		"cinema_id": cinemaId,
	})
}
