package reply

// Object 数据返回格式
type Object struct {
	State  int         `json:"state"`
	Object interface{} `json:"object"`
}

// Array 数据返回格式
type Array struct {
	State int         `json:"state"`
	Rows  interface{} `json:"rows"`
	Total int         `json:"total"`
}

// Error 错误返回格式
type Error struct {
	State   int    `json:"state"`
	Message string `json:"message"`
}

// Urlencoded 错误
func Urlencoded() *Error {
	return &Error{State: 1001, Message: "is not Urlencoded"}
}

// Multipart 错误
func Multipart() *Error {
	return &Error{State: 1002, Message: "is not Multipart"}
}
