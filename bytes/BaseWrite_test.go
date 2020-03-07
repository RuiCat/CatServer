package bytes

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	// 测试结构体
	type T struct {
		A int
		B string
		C map[int]string
		D []float32
		E interface{}
	}

	// 解析结构体
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: binary.LittleEndian}

	// 写入结构体
	fmt.Println("写入流: ", BaseWrite(w, T{
		A: 2233,
		B: "喵~",
		C: map[int]string{1: "喵~"},
		D: []float32{22.33},
		E: "喵呜~",
	}))
	fmt.Println("流数据:", w.Byte)

	// 读出结构体
	v := &T{}
	BaseRead(r, v)
	fmt.Println("还原流:", v)
}
