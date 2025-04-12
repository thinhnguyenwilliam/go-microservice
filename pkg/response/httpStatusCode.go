package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Email is invalid
	ErrorInvalidToken     = 30003 // Invalid token
)

var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorInvalidToken:     "Invalid token",
}

func Msg(code int) string {
	if message, ok := msg[code]; ok {
		return message
	}
	return "Unknown error"
}
