package conn

import (
	"CatServer/pool"
	"net"

	"CatServer/errors"
)

// server 服务器
type server struct {
	conn    chan NetConn
	pool    pool.Pool
	err     errors.Error
	close   func() error
	isclose bool
}

// Listen 监听
func (s *server) Listen(network, address string) (err error) {
	// 拦截错误
	defer errors.RecoverFunc(s.err, func() {
		err = s.err
	})
	// 判断协议类型
	switch network {
	case "tcp", "tcp4", "tcp6", "unix", "unixpacket":
		// 打开连接
		tcp, e := net.Listen(network, address)
		errors.Panic(CodesTCP, e)
		// 进入循环
		s.isclose = true
		go func() {
			var conn net.Conn
			for {
				// 得到缓存对象
				v, p := s.pool.Get()
				c, _ := v.(*Conn)
				c.Pool = p
				// 得到连接
				conn, c.Error = tcp.Accept()
				if !s.isclose {
					return
				}
				// 处理错误
				if c.Error == nil {
					// 绑定信息
					c.Addr = conn.RemoteAddr()
					c.Codes = CodesTCPConn
					c.callRead = conn.Read
					c.callClose = conn.Close
					c.callTCPWrite = conn.Write
					// 发送到处理
					s.conn <- c
				} else {
					s.err.AddCodes(CodesTCPConn, c.Error)
					c.Pool.Put() // 出错不进行处理
				}
			}
		}()
		s.close = tcp.Close
	case "udp", "udp4", "udp6":
		// 打开udp连接
		addr, e := net.ResolveUDPAddr(network, address)
		errors.Panic(CodesUDP, e)
		udp, e := net.ListenUDP(network, addr)
		errors.Panic(CodesUDP, e)
		// 发送处理
		s.isclose = true
		go func() {
			var n int
			for {
				// 得到缓存对象
				v, p := s.pool.Get()
				c, _ := v.(*Conn)
				c.Pool = p
				// 得到连接
				n, c.Addr, c.Error = udp.ReadFrom(c.Data)
				if !s.isclose {
					return
				}
				// 处理错误
				if c.Error == nil {
					// 绑定信息
					c.Data = c.Data[:n]
					c.Codes = CodesUDPConn
					c.callUDPWrite = udp.WriteTo
					// 发送到处理
					s.conn <- c
				} else {
					s.err.AddCodes(CodesTCPConn, c.Error)
					c.Pool.Put() // 出错不进行处理
				}
			}
		}()
		// 记录关闭
		s.close = udp.Close
	default:
		panic("异常协议")
	}
	return nil
}

// Accept 读取
func (s *server) Accept() (c NetConn, b bool) {
	c, b = <-s.conn
	return c, b
}

// Close 关闭
func (s *server) Close() error {
	s.isclose = false // 标记关闭
	s.close()         // 关闭服务
	close(s.conn)     // 关闭写通道
	// 关闭客户
	s.pool.Traversal(func(v *pool.Data) {
		(v.Value.(NetConn)).Close()
	})
	return s.err
}

// Init 初始化
//  n 对象缓存大小
//  b 连接数据缓冲大小
func Init(n int, b int) New {
	s := &server{}
	s.conn = make(chan NetConn, n)
	s.pool = pool.Pool{
		New: func() interface{} {
			return &Conn{Data: make([]byte, b)}
		},
	}
	s.pool.Init(n)
	s.err = errors.New(CodesInit, "服务器创建")
	return s
}
