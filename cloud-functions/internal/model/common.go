package model

type ApiResponse struct {
	Success bool   `json:"success"`        // 处理是否成功
	Code    int    `json:"code"`           // 状态码
	Error   string `json:"error"`          // 错误信息
	Data    any    `json:"data,omitempty"` // 结果数据
}

type BusinessError struct {
	Message string
	Code    int
}

func (e *BusinessError) Error() string {
	return e.Message
}

func NewBusinessError(message string, code int) *BusinessError {
	return &BusinessError{Message: message, Code: code}
}
