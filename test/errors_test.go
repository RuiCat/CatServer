package test

import (
	"CatServer/errors"
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	c := errors.Add(5)
	// 测试创建错误
	err := errors.New(c, fmt.Errorf("创建错误测试"))
	err.Warp(fmt.Errorf("%s", "继承了一个错误喵~"))
	fmt.Println(err.Error())
}
