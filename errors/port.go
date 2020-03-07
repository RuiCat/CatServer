package errors

import "fmt"

// Error 错误
type Error interface {
	SetCode(code Codes)
	AddCodes(c Codes, v error)
	Warp(err error)
	Get() error
	Error() string
	Is(c Codes, v error) bool
}

// Panic 弹出错误
func Panic(c Codes, v interface{}) {
	if v != nil {
		panic(&errors{code: c, path: caller(2), data: []error{fmt.Errorf("%s", v)}})
	}
}

// Recover 拦截错误
func Recover(e error) {
	if err := recover(); err != nil {
		if er, ok := e.(Error); ok {
			er.Warp(fmt.Errorf("%s", err))
		} else {
			// 如果无法继承错误则丢出去
			panic(err)
		}
	}
}

// RecoverFunc 拦截错误并回调
func RecoverFunc(e error, f func()) {
	if err := recover(); err != nil {
		if er, ok := e.(Error); ok {
			er.Warp(fmt.Errorf("%s", err))
			f()
		} else {
			// 如果无法继承错误则丢出去
			panic(err)
		}
	}
}

// New 创建
func New(c Codes, v interface{}) Error {
	return &errors{code: c, path: caller(2), data: []error{fmt.Errorf("%s", v)}}
}

// NewObj 创建
func NewObj(c Codes, obj interface{}) Error {
	return &errors{code: c, path: caller(2), data: []error{fmt.Errorf("%v", obj)}}
}
