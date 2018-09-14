package errors

import (
	"encoding/json"
	"fmt"
)

var (
	ApiErrRefuse           = ApiStandardError{403, "服务器拒绝请求"}
	ApiErrSuccess          = ApiStandardError{1000, "正常!请安心服用!"}
	ApiErrNamePwdIncorrect = ApiStandardError{1001, "用户名或密码错误"}
	ApiErrPwdIncorrect = ApiStandardError{1002, "密码错误"}
)

type ApiStandardError struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func NewUnknownErr(e error) ApiStandardError {
	return ApiStandardError{400, e.Error()}
}

func (err ApiStandardError) Error() string {
	bytes, e := json.Marshal(err)
	if e != nil {
		return NewUnknownErr(e).Error()
	}
	fmt.Println(string(bytes))
	return string(bytes)
}
