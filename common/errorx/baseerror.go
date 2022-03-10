package errorx

const defaultCode = 1001

type CodError struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func NewCodeError(code int,msg string)error {
	return &CodError{
		Code: code,
		Msg: msg,
	}
}

func  NewDefaultError(msg string)error  {
	return NewCodeError(defaultCode,msg)
}

func (e *CodError)Error()string  {
	return e.Msg
}

func (e *CodError)Data() *CodeErrorResponse  {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg: e.Msg,
	}
}