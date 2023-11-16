package main

import "fmt"

// 空接口和类型断言
func DoSomething(p interface{})  {
	/*if i, ok := p.(int); ok {
		fmt.Println("Integer : ", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string : ", i)
		return
	}
	fmt.Println("Unknow type")*/
	switch i:=p.(type) {
	case int:
		fmt.Println("Integer : ", i)
	case string:
		fmt.Println("string : ", i)
	default:
		fmt.Println("Unknow type")
	}
}

func main() {
	DoSomething(1)
	DoSomething("zhizhi")
	DoSomething(1.0)
}
