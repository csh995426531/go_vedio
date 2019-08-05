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

// Normal 正常
type Normal struct {
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

var (
	// ErrorRequestBodyParseFaild 请求内容解析失败
	ErrorRequestBodyParseFaild = ErrResponse{HttpSC: 400, Error: Err{Error: "request body parse fail", ErrorCode: "001"}}
	// ErrorNotAuthUser 请求用户认证失败
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "user  auth fail", ErrorCode: "002"}}
	// ErrorDBError DB失败
	ErrorDBError = ErrResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	// ErrorInternalFaults 内部故障
	ErrorInternalFaults = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
	// ErrorSessionError Session处理错误
	ErrorSessionError = ErrResponse{HttpSC: 500, Error: Err{Error: "Session handler error", ErrorCode: "005"}}
	// ErrorUserAlreadyExsitsError  用户已存在
	ErrorUserAlreadyExsitsError = ErrResponse{HttpSC: 403, Error: Err{Error: "User already exsits", ErrorCode: "006"}}
)
