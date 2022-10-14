package param

type Response struct {
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Payload      interface{} `json:"payload"`
}
