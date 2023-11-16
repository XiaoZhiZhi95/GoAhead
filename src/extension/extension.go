package main

import "fmt"

type Pet struct {

}

func (p *Pet)Speak() {
	fmt.Println("I am Pet")
}

type Dog struct {
	Pet
}

// 当Dog没有定义该方法，则d.Speak()会输出"I am Pet"
// Dog定义了一样的方法名，则d.Speak()会输出自己定义的"Wang wang"
/*func (d *Dog)Speak()  {
	fmt.Println("Wang wang")
}*/

// go语言的复合和扩展，没有继承
func main() {
	// 无法定义，会编译报错，因为不是继承关系
	/*var d *Pet
	d = new(Dog)*/
	d := &Dog{}
	d.Pet.Speak()
	d.Speak()
}
