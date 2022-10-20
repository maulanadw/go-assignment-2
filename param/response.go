package param

type Response struct {
	Status    int         `json:"status"`
	Message   *string     `json:"message,omitempty"`
	ErrorInfo *string     `json:"error_info,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
}
