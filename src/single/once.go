package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type Single struct {
	i int
}

var SingleInstance *Single
var once = sync.Once{}

func GetSingleInstance(i int) *Single {
	// once := sync.Once{}	// 单例模式once，必须是全局的，如果放在这，则不会是全局多线程间执行一次
	once.Do(func() {
		fmt.Println("create singleInstance")
		SingleInstance = &Single{i: i}
	})
	return SingleInstance
}

func main() {
	for i := 0; i<10;i++ {
		go func() {
			obj := GetSingleInstance(i)
			fmt.Printf("pointe = %d, obj=%v\n", unsafe.Pointer(obj), obj)	// 此处&obj取的是临时变量的地址，unsafe.Pointer将obj的值*Single指针转换成了一个通用的指针
		}()
	}
	time.Sleep(1*time.Second)
}
