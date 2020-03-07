package group

// 注释
//+ 在编写调动时做了各种实现但是均不满足整个系统对调度的需求.
//+ 因为不同的也业务资源消耗是不同的无法再调度时候对业务进行限制.
//+ 实现采用了
//+                             入口服务                               业务实现
//+   [UDP/TCP] => {Group -> 序列化 -> 转发到业务路由} => 业务路由 -> {Group => 业务限制} -> 业务实现

// Group 调度器
type Group struct {
	g chan func()
}

// Init 初始化
//  n 携程数量
func (g *Group) Init(n int) {
	g.g = make(chan func(), n)
	//+ 进行调度
	go func() {
		var f func()
		for f = range g.g {
			f()
		}
	}()
}

// Go 调度
//+ 执行成功返回真
func (g *Group) Go(f func()) bool {
	select {
	case g.g <- f:
	default:
		return false
	}
	return true
}

// GoLock 等待调度
//+ 如果无法调度则等待
func (g *Group) GoLock(f func()) {
	g.g <- f
}

// Close 关闭
//+ 非安全操作,因为比起在底层保证安全在代码实现时保证效率更高.
func (g *Group) Close() {
	close(g.g)
}
