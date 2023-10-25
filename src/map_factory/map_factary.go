package main

import "log"

func main() {
	m := make(map[string]func(a, b int) int)
	m["add"] = func(a, b int) int {
		return a + b
	}
	m["multi"] = func(a, b int) int {
		return a * b
	}

	log.Println("add = ", m["add"](1, 2))
	log.Println("multi = ", m["multi"](1, 2))
}
