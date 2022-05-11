package model

type Result struct {
	Code    string
	Message string
	Data    interface{}
}

func Success() Result {
	return Result{"200", "", nil}
}

func NewResult(code string, msg string, data interface{}) Result {
	return Result{Code: code, Message: msg, Data: data}
}

func (a Result) IsSuccess() bool {
	return a.Code == "200"
}

func (a Result) SetCode(code string) Result {
	a.Code = code
	return a
}

func (a Result) SetMsg(msg string) Result {
	a.Message = msg
	return a
}

func (a Result) SetData(data interface{}) Result {
	a.Data = data
	return a
}
