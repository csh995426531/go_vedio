package defs

// Err 错误
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// ErrResponse 错误响应
type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	// ErrorRequestBodyParseFaild 请求内容解析失败
	ErrorRequestBodyParseFaild = ErrResponse{HttpSC: 400, Error: Err{Error: "request body parse fail", ErrorCode: "001"}}
	// ErrorNotAuthUser 请求用户认证失败
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "user  auth fail", ErrorCode: "002"}}
)
