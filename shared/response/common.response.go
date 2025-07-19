package response

type SuccessfullyResponse struct {
	Message string      `json:"message"`
	Item    interface{} `json:"item,omitempty"`
	Items   interface{} `json:"items,omitempty"`
}
