package bytes

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBit64(t *testing.T) {
	fmt.Println("Int为64位:", bit64)
}

func TestWriter(t *testing.T) {
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: w.Order}
	// 写测试
	w.Int(2233)
	w.Bool(true)
	w.Bytes([]byte("喵喵喵~"))
	w.Float32(22.33)
	fmt.Println("写入结果:", w.Byte)
	// 读测试
	fmt.Println("读:", r.Int())
	fmt.Println("读:", r.Bool())
	fmt.Println("读:", string(r.Bytes()))
	fmt.Println("读:", r.Float32())
	fmt.Println("剩余:", r.Byte)
}

func TestTw(t *testing.T) {
	b := make([]byte, 0, 80000000)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	for i := 0; i < 10000000; i++ {
		w.Int(i)
	}
	fmt.Println("流大小:", len(*w.Byte))
}
