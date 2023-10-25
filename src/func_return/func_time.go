package main

import (
	"fmt"
	"time"
)

// 实现一个完整的功能，计算函数的运行时长
func main() {
	newFunc := timeSpent(slowFunc)	// 把一个函数包起来，形成另一个函数，函数式编程
	newFunc(10)
}

// 计算函数运行时间的通用方法，将需要进行计算的函数传进来，包装得到新函数
func timeSpent(fn func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := fn(n)
		fmt.Println("spent = ", time.Since(start).Seconds())
		return ret
	}
}

// 一个需要计算运行时间的函数本身
func slowFunc(op int) int {
	time.Sleep(1*time.Second)
	fmt.Println("slowFunc param = ", op)
	return op
}