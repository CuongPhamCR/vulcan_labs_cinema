package validation

import (
	"net/http"
	"vulcan_labs_cinema/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Hàm format lỗi validation
func FormatValidationError(err error) []string {
	errMessages := []string{}

	if err.Error() == "EOF" {
		errMessages = append(errMessages, "Body is empty")
		return errMessages
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range validationErrors {
			// fieldName := fe.Field()
			// append(errMessages, getErrorMessage(fe))

			// append to errMessages
			errMessages = append(errMessages, getErrorMessage(fe))
		}
	}
	return errMessages
}

// Hàm trả về thông điệp lỗi tùy chỉnh
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " không được để trống."
	case "email":
		return "Email không hợp lệ."
	case "min":
		return fe.Field() + " phải có ít nhất " + fe.Param() + " ký tự."
	case "max":
		return fe.Field() + " không được quá " + fe.Param() + " ký tự."
	case "gte":
		return fe.Field() + " phải lớn hơn hoặc bằng " + fe.Param() + "."
	case "lte":
		return fe.Field() + " phải nhỏ hơn hoặc bằng " + fe.Param() + "."
	case "len":
		return fe.Field() + " phải có đúng " + fe.Param() + " ký tự."
	default:
		return "Giá trị của " + fe.Field() + " không hợp lệ."
	}
}

// Middleware xử lý validation
func ValidationMiddleware(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(obj); err != nil {
			errMessages := FormatValidationError(err)
			if len(errMessages) > 0 {
				response.ErrorResponse(c, http.StatusBadRequest, response.ErrCodeParamInvalid, errMessages[0])
			} else {
				response.ErrorResponse(c, http.StatusBadRequest, response.ErrCodeParamInvalid, "Param is invalid")
			}
			c.Abort()
			return
		}
		c.Set("requestData", obj)
		c.Next()
	}
}
