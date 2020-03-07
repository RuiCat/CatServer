package main

import (
	"CatServer/bytes"
	"CatServer/conn"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// ProgramServer 服务
// port 端口
func ProgramServer(port string) error {
	// 启动读取服务器
	server := conn.Init(10, 255)
	// 启动服务器
	if err := server.Listen("udp", port); err != nil {
		return err
	}
	// 用户列表
	list := make(map[string]Program)
	go func() {
		tk := time.NewTicker(time.Minute)
		for range tk.C {
			fmt.Println("发送心跳包")
			for _, p := range list {
				// 测试发送心跳包
				udp, err := net.Dial("udp", p.Port)
				if err != nil {
					fmt.Println(err)
					return
				}
				udp.Write([]byte("喵~"))
			}
		}
	}()
	// 等待用户连接
	for {
		co, ok := server.Accept() // 得到链接
		if !ok {
			return nil
		}
		// 转交处理函数
		go func(co conn.NetConn) {
			// 处理用户中心发来请求
			b, _ := co.Read()
			r := &bytes.Read{Byte: &b, Order: binary.LittleEndian}
			// 得到用户ID
			p := Program{
				SSH:   string(r.Bytes()),
				Start: string(r.Bytes()),
				Addr:  string(r.Bytes()),
				Port:  string(r.Bytes()),
			}
			list[co.RemoteAddr().String()] = p
			fmt.Println("新注册链接: ", p)
		}(co)
	}
	return nil
}

func main() {
	fmt.Println(ProgramServer(":88"))
}
