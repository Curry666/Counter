package c

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Mp sync.Map
}

func (c *Counter) Init() {
	*c = Counter{sync.Map{}}
}

func (c *Counter) Incr(key interface{}, value int) {
	if origin, ok := c.Mp.LoadOrStore(key, value); ok {
		c.Mp.Store(key, origin.(int)+value)
	}
}
func (c *Counter) Reset() {
	fmt.Println("重置计数器")
	c.Init()
}

//单位是ms
//间隔时间，执行函数，总执行时间
func (c *Counter) Flush2Broker(interval int, FuncCbFlush func(), total int) {
	if total < interval {
		fmt.Println("总时间小于间隔时间")
		return
	}
	go func() {
		ticker := time.NewTicker(time.Millisecond * time.Duration(interval))
		go func() {
			for {
				<-ticker.C
				FuncCbFlush()
				c.Reset()
			}
		}()
		//总共执行 total ms
		time.Sleep(time.Millisecond * time.Duration(total))
		ticker.Stop()
	}()
}
