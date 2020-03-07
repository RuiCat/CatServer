package errors

import (
	"fmt"
	"runtime"
)

// Codes 错误代码
type Codes uint32

// _codes 错误代码列表
var _codes = map[Codes]struct{}{}

// Add 添加一个错误代码
func Add(sign Codes) Codes {
	if _, ok := _codes[sign]; ok {
		panic(fmt.Sprintf("错误代码: %d 已经存在", sign))
	}
	_codes[sign] = struct{}{}
	return sign
}

// 默认定义
var (
	CodesNil = Add(0) // 零值
)

// caller 得到来源
func caller(id int) string {
	_, file, line, _ := runtime.Caller(id)
	return fmt.Sprintf("%s:%d", file, line)
}
