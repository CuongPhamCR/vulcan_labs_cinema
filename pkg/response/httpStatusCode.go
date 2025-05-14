package response

const (
	ErrCodeSuccess        = 20001
	ErrCodeInternalServer = 20002
	ErrCodeParamInvalid   = 20003
	ErrCodeInvalidCount   = 20004
	ErrCodeCinemaNotFound = 20005
)

// Message
var msg = map[int]string{
	ErrCodeSuccess:        "Success",
	ErrCodeInternalServer: "Internal server error",
	ErrCodeParamInvalid:   "Param is invalid",
	ErrCodeInvalidCount:   "Invalid count",
	ErrCodeCinemaNotFound: "Cinema not found or not initialized",
}
