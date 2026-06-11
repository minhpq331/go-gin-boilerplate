package common

type responseAPI[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func SendResponse[T any](data T) *responseAPI[T] {
	return &responseAPI[T]{
		Data: data,
	}
}

func SendError(apiErr *APIError) *responseAPI[any] {
	return &responseAPI[any]{
		Code:    apiErr.ErrorCode,
		Message: apiErr.Message,
	}
}
