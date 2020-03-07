package conn

import (
	"CatServer/errors"
	"CatServer/pool"
	"io"
	"net"
)

// Conn 用户连接
type Conn struct {
	Data  []byte
	Pool  *pool.Data
	Addr  net.Addr
	Error error
	Codes errors.Codes
	// 回调事件
	callClose    func() error
	callRead     func(b []byte) (n int, err error)
	callTCPWrite func(b []byte) (int, error)
	callUDPWrite func(b []byte, addr net.Addr) (int, error)
}

// Write 写
func (c *Conn) Write(b []byte) (int, error) {
	switch c.Codes {
	case CodesTCPConn:
		return c.callTCPWrite(b)
	case CodesUDPConn:
		return c.callUDPWrite(b, c.Addr)
	}
	return 0, errors.NewObj(CodesConnWrite, c)
}

// Read 读数据
func (c *Conn) Read() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}
	// 处理读
	switch c.Codes {
	case CodesTCPConn:
		n, err := c.callRead(c.Data)
		c.Data = c.Data[:n]
		return c.Data, err
	case CodesUDPConn:
		// 标记一次读如果在读的话就报错退出
		//+ udp的 conn 读是一次性的如果读完则退出不对来源处理.
		c.Error = io.EOF
		return c.Data, nil
	}
	return c.Data, errors.NewObj(CodesConnRead, c)
}

// Close 关闭
func (c *Conn) Close() error {
	c.Pool.Put()
	if c.Codes == CodesTCPConn {
		return c.callClose()
	}
	return nil
}

// RemoteAddr 远程连接
func (c *Conn) RemoteAddr() net.Addr {
	return c.Addr
}
