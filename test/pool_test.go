package test

import (
	"CatServer/pool"
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	p := pool.Pool{
		New: func() interface{} {
			return "喵喵喵喵~"
		},
	}
	p.Init(100)

	s, v := p.Get()
	v.Put()
	fmt.Println(s)
}

func TestPool_T(t *testing.T) {
	p := pool.Pool{
		New: func() interface{} {
			return "喵喵喵喵~"
		},
	}
	p.Init(100)

	for i1 := 0; i1 < 10; i1++ {
		t1 := *pool.Ts
		for i := 0; i < 100000; i++ {
			_, v := p.Get()
			//<-time.NewTicker(time.Second * 2).C
			v.Put()
		}
		fmt.Println((*pool.Ts).Sub(t1))
	}
	fmt.Println(p.Tlist)
}
