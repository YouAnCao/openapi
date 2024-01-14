package common

import (
	"errors"
	"fmt"
)

// OpenapiError 平台异常
type OpenapiError struct {
	Code int
	error
}

// BusinessError 业务异常
type BusinessError struct {
	OpenapiError
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.error.Error())
}

// NewError 新异常
func NewError(code int, msg string) BusinessError {
	return BusinessError{
		OpenapiError: OpenapiError{
			Code:  code,
			error: errors.New(msg),
		},
	}
}

// AccountNotFound 账户不存在
var accountNotFound = NewError(1, "账户不存在!")

// TokenNotValidate token验证失败
var tokenNotValidate = NewError(2, "token 验证失败")
