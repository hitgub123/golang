package obj

type Result struct {
	Code int8
	Error error
	ErrorMessage string
	Data interface{}
}