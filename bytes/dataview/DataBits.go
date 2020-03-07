package dataview

// DataBits 字节
type DataBits [8]bool

// Set 设置
func (db *DataBits) Set(b byte) {
	db[0] = ((b >> 0) & 0x00000001) == 1
	db[1] = ((b >> 1) & 0x00000001) == 1
	db[2] = ((b >> 2) & 0x00000001) == 1
	db[3] = ((b >> 3) & 0x00000001) == 1
	db[4] = ((b >> 4) & 0x00000001) == 1
	db[5] = ((b >> 5) & 0x00000001) == 1
	db[6] = ((b >> 6) & 0x00000001) == 1
	db[7] = ((b >> 7) & 0x00000001) == 1
}

// Get 读取
func (db *DataBits) Get() byte {
	var b uint8
	if db[0] {
		b = b | 1<<0
	}
	if db[1] {
		b = b | 1<<1
	}
	if db[2] {
		b = b | 1<<2
	}
	if db[3] {
		b = b | 1<<3
	}
	if db[4] {
		b = b | 1<<4
	}
	if db[5] {
		b = b | 1<<5
	}
	if db[6] {
		b = b | 1<<6
	}
	if db[7] {
		b = b | 1<<7
	}
	return b
}
