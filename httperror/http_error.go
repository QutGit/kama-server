package httperror

import "fmt"

// HTTPError 自定义HTTP 错误
type HTTPError struct {
	// Http 状态码
	StatusCode int `json:"http_code"`
	// 错误消息
	Message string `json:"msg"`
	// 业务错误码
	Code int64 `json:"code"`
	// IsHttpError 标记
	IsHTTPError bool `json:"-"`
}

func (that *HTTPError) Error() error {
	return fmt.Errorf("%d-%s", that.Code, that.Message)
}

// IsHTTPError 判断是否是 HTTPError 类型
func IsHTTPError(e interface{}) bool {
	v, ok := e.(*HTTPError)
	fmt.Println(v)
	return ok
}

// New 生成 HTTPError
func New(statusCode int, message string, code int64) *HTTPError {
	if statusCode > 600 {
		panic("statusCode不能大于600")
	}
	return &HTTPError{
		StatusCode:  statusCode,
		Message:     message,
		Code:        code,
		IsHTTPError: true,
	}
}

// BadRequest 400 参数校验失败
func BadRequest(message string, code int64) *HTTPError {
	return New(400, message, code)
}

// InternalError 500 系统内部错误
func InternalError(message string, code int64) *HTTPError {
	return New(500, message, code)
}

// RequestNotAccept 406 请求不可接受
func RequestNotAccept(message string, code int64) *HTTPError {
	return New(406, message, code)
}
