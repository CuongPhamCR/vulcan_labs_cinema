package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, httpCode int, customCode int, message string) {
	if httpCode == 0 {
		httpCode = http.StatusInternalServerError
	}
	c.JSON(httpCode, ResponseData{
		Code: customCode,
		Message: func() string {
			if message != "" {
				return message
			}
			return msg[customCode]
		}(),
		Data: nil,
	})
}

func GetErrorMessage(code int) string {
	return msg[code]
}
