package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Email is invalid
)

var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Email is invalid",
}
