package errors

import "fmt"

// errors 错误信息
type errors struct {
	code Codes
	path string
	data []error
}

// SetCode 设置错误代码
func (e *errors) SetCode(code Codes) {
	e.code = code
}

// AddCodes 添加一个错误代码
func (e *errors) AddCodes(c Codes, v error) {
	e.data = append(e.data, &errors{code: c, path: caller(2), data: []error{v}})
}

// Warp 添加一个错误
func (e *errors) Warp(err error) {
	if err != nil && err.Error() != "" {
		e.data = append(e.data, err)
	}
}

// Get 得到错误
func (e *errors) Get() error {
	if len(e.data) == 0 {
		return nil
	}
	return e
}

// Error 得到错误信息
func (e *errors) Error() string {
	if len(e.data) == 0 {
		return ""
	}
	// 构建错误
	s := e.data[0].Error()
	for _, st := range e.data[1:] {
		s += "\n::" + st.Error()
	}
	return fmt.Sprintf("[%d][%s]: %s", e.code, e.path, s)
}

// Is 错误判定用
func (e *errors) Is(c Codes, v error) bool {
	if v != nil && v.Error() != "" {
		e.AddCodes(c, v)
		return true
	}
	return false
}
