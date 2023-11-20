package main

import (
	"errors"
	"fmt"
	"time"
)

// 复杂对象的结构
type ReusableObj struct {
	Sequence int	// 对象序号
}

func (obj *ReusableObj)String() {
	fmt.Println("This obj Sequence is ", obj.Sequence)
}

// 一个对象池，包含这个对象的有buffer的channel
type ObjPool struct {
	Pool chan *ReusableObj
}

// NewReusablePool 初始化对象池
func NewReusablePool(num int) *ObjPool {
	pool := ObjPool{}
	pool.Pool = make(chan *ReusableObj, num)
	for i:=0; i<num; i++ {
		pool.Pool <- &ReusableObj{Sequence: i}
	}

	return &pool
}

// 获取对象，支持超时时间控制
func (p *ObjPool)GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.Pool:
		return ret, nil
	case <- time.After(timeout):	// 超时控制
		return nil, errors.New(fmt.Sprintf("Get ReusableObj Time Out, time = %v", timeout))
	}
}

// 归还对象到对象池
func (p *ObjPool)ReleaseObj(obj *ReusableObj) error {
	select {
	case p.Pool <- obj:
		return nil
	default:	// 池子满了，无法及时归还，则可以直接返回
		return errors.New("Release Obj err with overflow")
	}
}

func main()  {
	num := 10
	pool := NewReusablePool(num)
	for i := 0; i < num+1; i++ {
		obj, err := pool.GetObj(1*time.Second) 	// 获取对象
		if err != nil {
			fmt.Println("get obj get err = ", err)
		} else {
			obj.String()
			// 释放对象
			err = pool.ReleaseObj(obj)
			if err != nil {
				fmt.Println("release obj get err = ", err)
			}
		}
	}
}