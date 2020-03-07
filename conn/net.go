package conn

import (
	"CatServer/errors"
	"net"
)

// 错误定义
var (
	CodesInit      = errors.Add(1) // 链接初始化错误
	CodesTCP       = errors.Add(2) // TCP 服务器
	CodesUDP       = errors.Add(3) // UDP 服务器
	CodesUDPConn   = errors.Add(4) // UDP 用户连接错误
	CodesTCPConn   = errors.Add(5) // TCP 用户连接错误
	CodesConnWrite = errors.Add(6) // Conn 写错误
	CodesConnRead  = errors.Add(7) // Conn 读错误
)

// New 网络
type New interface {
	Listen(network, address string) (err error)
	Accept() (NetConn, bool)
	Close() error
}

// NetConn 用户接口
type NetConn interface {
	Read() ([]byte, error)
	Write(b []byte) (n int, err error)
	Close() error
	RemoteAddr() net.Addr
}
