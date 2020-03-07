package test

import (
	"CatServer/group"
	"CatServer/pool"

	"fmt"
	"sync"
	"testing"
)

func TestGroup(t *testing.T) {
	// 测试
	f := group.Group{}
	f.Init(100)
	// 执行记录
	wait := sync.WaitGroup{}
	t1 := *pool.Ts
	for i := 0; i < 1000000; i++ {
		wait.Add(1)
		// 指针传值需要创建值对象
		f.GoLock(func() {
			wait.Done()
		})
	}
	wait.Wait()
	fmt.Println((*pool.Ts).Sub(t1))
}
