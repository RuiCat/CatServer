package test

import (
	"CatServer/conn"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	c := conn.Init(100, 255)
	fmt.Println(c.Listen("tcp", ":80"))
	// 循环得到连接
	for {
		co, ok := c.Accept()
		if !ok {
			return
		}
		go func(co conn.NetConn) {
			b, _ := co.Read()
			co.Write(b)
			co.Close()
		}(co)
	}
}
func TestDial(t *testing.T) {
	for index := 0; index < 100; index++ {
		go func(i int) {
			c, err := net.Dial("tcp", "127.0.0.1:80")
			if err != nil {
				fmt.Println(err)
				return
			}
			c.Write([]byte(fmt.Sprint(i)))
			b := make([]byte, 255)
			n, _ := c.Read(b)
			fmt.Println(string(b[:n]))
		}(index)
	}
	<-time.NewTicker(5 * time.Second).C
}
