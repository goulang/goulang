package errors

import "fmt"

var (
	ApiErrNamePwdIncorrect = ApiStandardError{1001, "用户名或密码错误"}
)

type ApiStandardError struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func (err ApiStandardError) Error() string {
	return fmt.Sprintf("code: %d, errorMsg %s", err.ErrorCode, err.ErrorMsg)
}
