package pool

import "time"

// Ts 时间
//  用于取时间差等时间上的纪录操作
var Ts = func() *time.Time {
	t := new(time.Time)
	*t = time.Now().UTC()
	go func() {
		tk := time.NewTicker(time.Microsecond)
		// 更新时间戳
		for range tk.C {
			*t = time.Now().UTC()
		}
	}()
	return t
}()
