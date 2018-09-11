package errors

import (
	"encoding/json"
)

type ApiError struct {
	HttpCode int
	Messages map[string]interface{}
}

func NewApiError(code int, errorMsg string) *ApiError {
	content := map[string]interface{}{
		"code":    code,
		"message": errorMsg,
	}

	return &ApiError{
		HttpCode: code,
		Messages: content,
	}
}

func (e *ApiError) Error() string {
	bytes, _ := json.Marshal(e.Messages)
	return string(bytes)
}
