package main

import (
	"CatServer/bytes"
	"CatServer/conn"
	"encoding/binary"
	"fmt"
	"net"
)

// Program 程序对象
type Program struct {
	SSH   string // 命令行链接(ssh)
	Start string // 程序启动(命令行)
	Addr  string // 连接地址(UDP)
	Port  string // 连接端口(UDP)
}

// Enroll 注册到维护中心
//  addr 远程地址
//  port 本地端口
func Enroll(
	addr, port string, // 软件信息
	ssh, start string, // 命令行信息
) (Program, error) {
	// 结构对象
	p := Program{Port: port, SSH: ssh, Start: start}
	// 连接维护中心
	udp, err := net.Dial("udp", addr)
	if err != nil {
		return p, err
	}
	// 程序内网IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return p, err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				p.Addr = ipnet.IP.String()
			}
		}
	}
	// 启动读取服务器
	server := conn.Init(10, 255)
	// 启动服务器
	if err := server.Listen("udp", port); err != nil {
		return p, err
	}
	// 构建连接信息
	w := &bytes.Write{Byte: new([]byte), Order: binary.LittleEndian}
	w.Bytes([]byte(p.SSH))
	w.Bytes([]byte(p.Start))
	w.Bytes([]byte(p.Addr))
	w.Bytes([]byte(p.Port))
	udp.Write(*w.Byte) // 发送
	// 循环读取心跳包
	for {
		co, ok := server.Accept() // 得到链接
		if !ok {
			return p, nil
		}
		// 转交处理函数
		go func(co conn.NetConn) {
			// 是否由管理中心发送来的请求
			fmt.Println(co.RemoteAddr().String())
			//if co.RemoteAddr().String() != addr {
			//	return
			//}
			// 处理用户中心发来请求
			b, _ := co.Read()
			fmt.Println("维护心发送的数据:", string(b))
		}(co)
	}
}
