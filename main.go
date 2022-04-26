package main

import (
	"fmt"
	"time"

	"./c"
)

func main() {
	Test()
}
func Test() {
	counter := &c.Counter{}

	counter.Init()
	counter.Incr("get.called", 123)
	value_1, _ := counter.Mp.Load("get.called")
	counter.Incr("get.called", 456)
	value_2, _ := counter.Mp.Load("get.called")
	fmt.Println(value_1, value_2)
	//刷新
	TestFunc := func() {
		fmt.Println("执行测试方法")
	}
	counter.Flush2Broker(5000, TestFunc, 20000)
	time.Sleep(20000 * time.Millisecond)
}
