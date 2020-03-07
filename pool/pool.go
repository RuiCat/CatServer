package pool

import (
	"sync"
)

// T 超时控制
type T struct {
	int64
}

// Data 池对象
type Data struct {
	Value interface{} // 值
	t     int64       // 开始时间
	look  bool        // 标记是否可用
	pool  *Pool       // 缓存池
}

// Put 放入
func (pd *Data) Put() {
	pd.look = false
	pd.pool.put(pd)
}

// Pool 缓冲池
type Pool struct {
	New    func() interface{} // 用于创建对象
	plist  map[*T]*Data       // 对象列表
	pcache []*Data            // 缓存列表
	pchan  chan *Data         // 缓冲对象通道
	Tlist  []int64            // 生存时间间隔表
	mutex  sync.Mutex
}

// Init 初始化
func (p *Pool) Init(n int) {
	p.pchan = make(chan *Data, n)
	p.plist = make(map[*T]*Data, n)
	// 初始化
	var v *Data
	for i := 0; i < n; i++ {
		v = &Data{Value: p.New(), pool: p}
		p.plist[&T{-1}] = v
		p.pchan <- v
	}
}

// GetTlist 得到时间间隔列表
func (p *Pool) GetTlist() (t []int64) {
	t = make([]int64, len(p.Tlist))
	p.mutex.Lock()
	copy(t, p.Tlist)
	p.Tlist = p.Tlist[0:0]
	p.mutex.Unlock()
	return t
}

// Len 缓存大小
func (p *Pool) Len() int {
	return len(p.pcache)
}

// Cap 总占用大小
func (p *Pool) Cap() int {
	return len(p.plist)
}

// put 放入一个对象
func (p *Pool) put(v *Data) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	v.look = false
	select {
	case p.pchan <- v:
	default:
		// 将无法发送到缓冲通道的对象添加到缓冲列表
		p.pcache = append(p.pcache, v)
	}
	// 将时间添加到执行时间列表
	t := (Ts).UnixNano() - v.t
	// 如果执行时间不大于1s则不添加到超时列表
	if t > 0 {
		p.Tlist = append(p.Tlist, t)
	}
}

// Get 取出一个对象
func (p *Pool) Get() (_ interface{}, v *Data) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// 得到对象
	select {
	case v = <-p.pchan: // 得到对象
	default: // 无法获得
		// 判断缓冲列表是否有值
		if len(p.pcache) > 0 {
			// 通过缓冲得到值
			v = p.pcache[0]
			p.pcache = p.pcache[1:len(p.pcache)]
		} else {
			v = &Data{Value: p.New(), pool: p}
			p.plist[&T{Ts.UnixNano()}] = v
		}
	}
	// 设置值
	v.t = (Ts).UnixNano()
	v.look = true
	return v.Value, v
}

// Traversal 历遍
func (p *Pool) Traversal(f func(v *Data)) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, d := range p.plist {
		if d.look {
			f(d)
		}
	}
}
