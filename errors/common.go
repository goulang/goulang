package errors

import (
	"encoding/json"
	"fmt"
)

var (
	ApiErrRefuse           = ApiStandardError{403, "服务器拒绝请求"}
	ApiErrNamePwdIncorrect = ApiStandardError{1001, "用户名或密码错误"}
)

type ApiStandardError struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func NewUnknownErr(e error) ApiStandardError {
	return ApiStandardError{500, e.Error()}
}

func (err ApiStandardError) Error() string {
	bytes, e := json.Marshal(err)
	if e != nil {
		return NewUnknownErr(e).Error()
	}
	fmt.Println(string(bytes))
	return string(bytes)
}
