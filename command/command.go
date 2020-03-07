package command

import (
	"bufio"
	"io"
	"os/exec"
)

// Command 命令行
type Command interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Error(p []byte) (n int, err error)
	Exit() error
}

// New 创建终端
func New(s string) (_ Command, err error) {
	// 创建对象
	C := &cmd{}
	C.Cmd = exec.Command(s)
	// 标准输入
	C.Stdin, err = C.Cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	// 标准输出
	Stdout, err := C.Cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	// 标准错误
	Stderr, err := C.Cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	// 输入输出变量
	C.Stdout = bufio.NewReader(Stdout)
	C.Stderr = bufio.NewReader(Stderr)
	// 启动
	go C.Cmd.Run()
	return C, nil
}

// cmd 命令行
type cmd struct {
	Cmd    *exec.Cmd
	Stdin  io.WriteCloser // 标准输入
	Stdout *bufio.Reader  // 标准输出
	Stderr *bufio.Reader  // 标准错误
}

// Write 写入通道
func (C *cmd) Write(p []byte) (n int, err error) {
	return C.Stdin.Write(p)
}

// Read 读取通道
func (C *cmd) Read(p []byte) (n int, err error) {
	return C.Stdout.Read(p)
}

// Error 获取错误
func (C *cmd) Error(p []byte) (n int, err error) {
	return C.Stderr.Read(p)
}

// Exit 执行命令
func (C *cmd) Exit() error {
	if C.Cmd.Process == nil {
		return nil
	}
	return C.Cmd.Process.Kill()
}
