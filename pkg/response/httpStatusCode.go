package response

const (
	ErrCodeSuccess                   = 20001
	ErrCodeInternalServer            = 20002
	ErrCodeParamInvalid              = 20003
	ErrCodeInvalidCount              = 20004
	ErrCodeCinemaNotFound            = 20005
	ErrCodeSeatNotAvailableOrInvalid = 20006
	ErrCodeSeatNotFound              = 20007
	ErrCodeSeatIsNotBooked           = 20008
)

// Message
var msg = map[int]string{
	ErrCodeSuccess:                   "Success",
	ErrCodeInternalServer:            "Internal server error",
	ErrCodeParamInvalid:              "Param is invalid",
	ErrCodeInvalidCount:              "Invalid count",
	ErrCodeCinemaNotFound:            "Cinema not found or not initialized",
	ErrCodeSeatNotAvailableOrInvalid: "Seat not available or violates distance",
	ErrCodeSeatNotFound:              "Seat not found or not initialized",
	ErrCodeSeatIsNotBooked:           "Seat is not booked",
}
