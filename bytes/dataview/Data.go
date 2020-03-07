package dataview

import (
	"io"
)

// Data 数据结构体
type Data struct {
	Data   []byte // 数据
	Offset int    // 读取偏移
}

// Add 添加数据
func (d *Data) Add(b ...byte) {
	d.Data = append(d.Data, b...)
}

// Set 设置数据
func (d *Data) Set(offest int, b ...byte) error {
	if offest > len(d.Data) || offest < 0 {
		return io.EOF
	}
	// 计算偏移
	d.Offset = offest + len(b)
	if d.Offset > len(d.Data) {
		d.Data = append(d.Data, make([]byte, d.Offset-len(d.Data))...)
	}
	// 写入
	copy(d.Data[offest:], b)
	return nil
}

// Get 得到数据
func (d *Data) Get(offest int, size int) (b []byte) {
	if offest < 0 || offest > len(d.Data) {
		offest = d.Offset
	}
	// 拷贝数据
	n := offest + size
	if n > len(d.Data) {
		d.Offset = len(d.Data)
		b = append(b, d.Data[offest:]...)
	} else {
		d.Offset = n
		b = append(b, d.Data[offest:n]...)
	}
	return
}

// Read 读取数据
func (d *Data) Read(p []byte) (_ int, err error) {
	n := d.Offset      // 开始位置
	d.Offset += len(p) // 结束位置
	// 限制大小
	if d.Offset > len(d.Data) {
		err = io.EOF
		d.Offset = len(d.Data)
	}
	// 读取
	copy(p, d.Data[n:d.Offset])
	// 读取长度
	return d.Offset - n, err
}

// Write 写入数据
func (d *Data) Write(p []byte) (n int, err error) {
	d.Data = append(d.Data, p...)
	n = len(p)
	err = nil
	return
}

// Empty 清空
func (d *Data) Empty() (b []byte) {
	// 拷贝切片
	b = make([]byte, len(d.Data))
	copy(b, d.Data[:len(d.Data)])
	// 清空切片
	d.Data = d.Data[0:0]
	d.Offset = 0
	return
}
