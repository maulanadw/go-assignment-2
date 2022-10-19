package param

type Response struct {
	Status       int         `json:"status"`
	ErrorMessage *string     `json:"error_message,omitempty"`
	Payload      interface{} `json:"payload,omitempty"`
}
