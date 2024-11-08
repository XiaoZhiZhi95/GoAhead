package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	ret := map[int]bool {
		1: true,
		2: true,
	}
	delete(ret, 1)
	fmt.Println("len = ", len(ret))
}
