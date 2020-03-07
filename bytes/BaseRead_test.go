package bytes

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: w.Order}

	// 构建测试数据
	//- 设置 map 类型
	w.Int(2) // 数量
	//- 设置值
	w.Int(2233)             // key
	w.Bytes([]byte("喵喵喵~")) // value
	w.Int(3322)             // key
	w.Bytes([]byte("喵呜~~")) // value
	fmt.Println("字节流: ", w.Byte)

	// 构建 map
	vmap := make(map[int]string)
	fmt.Println(BaseRead(r, vmap))
	// 结果
	fmt.Println(vmap)
}

func TestSlice(t *testing.T) {
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: w.Order}

	// 测试数据
	// 长度
	w.Int(2)
	//+ 接口类型需要设置其类型
	w.Uint(uint(reflect.Int))
	w.Int(1)
	w.Uint(uint(reflect.String))
	w.Bytes([]byte("喵喵喵~"))

	// 读取数据
	var i []interface{}
	fmt.Println(BaseRead(r, &i))
	fmt.Println(i)
}
func TestArray(t *testing.T) {
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: w.Order}

	// 测试数据
	var a []int
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	// 数据
	for i, n := 0, len(b); i < n; i++ {
		w.Int(a[i])
	}

	// 读取数据
	var i [2][4]int
	fmt.Println(BaseRead(r, &i))
	fmt.Println(i)
}

func TestStruct(t *testing.T) {
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: w.Order}

	// 写入测试
	w.Bytes([]byte("喵喵喵~"))
	w.Bytes([]byte("喵呜~~"))

	// 测试字段
	s := struct {
		A string
		B string
	}{"", ""}
	fmt.Println(BaseRead(r, &s))
	fmt.Println(s.A, s.B)
}

func Test动态构建Map(t *testing.T) {
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: w.Order}

	// 写入测试
	//+ 设置类型
	w.Uint(uint(reflect.Map)) // Map类型
	//-> 创建map类型
	w.Uint(uint(reflect.Int)) // Key   类型
	w.Uint(uint(reflect.Int)) // value 类型
	//+ 设置值
	w.Int(2) // map数量

	w.Int(1)    // key
	w.Int(2233) // value

	w.Int(2)    // key
	w.Int(3322) // value

	// 测试字段
	fmt.Println("流:", w.Byte)
	var i interface{}
	fmt.Println("解析:", BaseRead(r, &i))
	fmt.Println("内容:", i.(*map[int]int))
}
func Test动态构建Array(t *testing.T) {
	b := make([]byte, 0)
	w := &Write{Byte: &b, Order: binary.LittleEndian}
	r := &Read{Byte: &b, Order: w.Order}

	// 写入测试
	//+ 设置类型
	w.Uint(uint(reflect.Array)) // Array类型与Slice类型相同
	w.Uint(uint(reflect.Int))   // 参数类型
	//-> 参数数量
	w.Int(2)
	//-> 写入值
	w.Int(0)
	w.Int(10)

	// 测试字段
	var i interface{}
	fmt.Println(BaseRead(r, &i))
	fmt.Println(i)
}
