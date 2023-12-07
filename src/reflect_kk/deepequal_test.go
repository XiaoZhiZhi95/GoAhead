package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T)  {
	a := map[int]string{1: "one", 2:"two", 3:"three"}
	b := map[int]string{1: "one", 2:"two", 3:"three"}

	// fmt.Println(a == b) // 不能进行比较，语法会报错
	fmt.Println(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 3, 2}

	fmt.Println("s1 == s2 ? ", reflect.DeepEqual(s1, s2))
	fmt.Println("s1 == s3 ? ", reflect.DeepEqual(s1, s3))

	c1 := MyCat{Name: "kk", Age:  5}
	c2 := MyCat{Name: "kk", Age: 5}
	fmt.Println("c1 == c2 ? ", reflect.DeepEqual(c1, c2))
}
