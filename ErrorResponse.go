package twitter

// ErrorResponse stores general Ridder API error response
//
type ErrorResponse struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Type   string `json:"type"`
}
