package response

const (
	ErrCodeSuccess        = 20001
	ErrCodeInternalServer = 20002
	ErrCodeParamInvalid   = 20003
)

// Message
var msg = map[int]string{
	ErrCodeSuccess:        "Success",
	ErrCodeInternalServer: "Internal server error",
	ErrCodeParamInvalid:   "Param is invalid",
}
