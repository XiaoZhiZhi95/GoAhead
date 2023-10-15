package main

import "log"

func main() {
	// 切片是可以共享的，切片间不可以进行 == 的比较
	slice := []int{0,1,2,3,4,5,6,7,8,9}
	slice1 := slice[2:4]
	log.Println("slice1 begin = ", slice1, ", len = ", len(slice1), ", cap = ", cap(slice1))
	slice2 := slice[2:5]
	log.Println("slice2 begin = ", slice2, ", len = ", len(slice2), ", cap = ", cap(slice2))
	log.Println("slice begin = ", slice)

	slice1[0] = 100
	log.Println("slice1 begin = ", slice1)
	log.Println("slice2 begin = ", slice2)
	log.Println("slice begin = ", slice)
}
