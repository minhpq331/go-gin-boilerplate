package common

type APIError struct {
	Message     string              `json:"message"`
	ErrorCode   int                 `json:"error_code"`
	Status      int                 `json:"status"`
	Stack       string              `json:"stack"`
	Errors      map[string][]string `json:"errors"`
	MessageData map[string]any      `json:"message_data,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(message string, errorCode int, status int, stack string) *APIError {
	return &APIError{
		Message:   message,
		ErrorCode: errorCode,
		Status:    status,
		Stack:     stack,
	}
}
