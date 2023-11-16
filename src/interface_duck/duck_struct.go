package main

import (
	"fmt"
)

type Base interface {
	WriteHelloWorld() string
}

type MyProgram struct {
	Name string
}

func (my *MyProgram)WriteHelloWorld() string {
	return fmt.Sprintf("Hello world, my name is %s", my.Name)
}

type MyWork struct {
	Status string
}

func (w *MyWork)WriteHelloWorld() string{
	return fmt.Sprintf("My work say %s to world", w.Status)
}

func HelloWorld(b Base) {
	fmt.Println(b.WriteHelloWorld())
}

func main() {
	/*var base Base
	base = &MyProgram{Name: "kk"}
	fmt.Println(base.WriteHelloWorld())

	base = &MyWork{Status: "happy"}
	fmt.Println(base.WriteHelloWorld())*/
	p := &MyProgram{Name: "xixi"}
	w := &MyWork{Status: "OK"}
	HelloWorld(p)
	HelloWorld(w)
}
