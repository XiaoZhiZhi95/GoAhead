package main

import (
	"fmt"
	"runtime"
	"sync"
)

var Pool = sync.Pool{
	New: func () interface{} {
		fmt.Println("sync pool create new")
		return 0
	},
}

func main()  {
	fmt.Println(Pool.Get())
	Pool.Put(100)
	fmt.Println(Pool.Get())
	fmt.Println(Pool.Get())

	Pool.Put("a")
	runtime.GC()	// GC 会清除sync.Pool中的缓存对象，将a清楚
	fmt.Println(Pool.Get())	// 此处会create new，而不是上面放入的a
}
