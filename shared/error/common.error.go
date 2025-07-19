package error

const (
	ERR_CONFLICT              string = "CONFLICT"
	ERR_BAD_REQUEST           string = "BAD_REQUEST"
	ERR_INTERNAL_SERVER_ERROR string = "INTERNAL_SERVER_ERROR"
	ERR_NOT_FOUND             string = "NOT_FOUND"
)

type ErrorResponse struct {
	Error ErrorDetailResponse `json:"error"`
}

type ErrorDetailResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
