package defs

type Err struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	//请求不正确
	ErrorRequestBodyParseFailed = ErrorResponse{400, Err{"Request Body is not correct", "001"}}

	//验证不通过，用户没有权限，用户不存在
	ErrorNotAuthorUser = ErrorResponse{401, Err{"User authentication failed", "002"}}

	//数据库操作执行出错
	ErrorDBExec = ErrorResponse{500, Err{"DB ops failure", "003"}}

	//序列化出错  对象 -- json
	ErrorInternalFaults = ErrorResponse{500, Err{"Internal service error", "004"}}
)
