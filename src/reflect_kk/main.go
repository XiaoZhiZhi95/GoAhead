package main

import (
	"fmt"
	"reflect"
)

func checkType(v interface{})  {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int16, reflect.Int64, reflect.Int8, reflect.Int32, reflect.Int:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func checkTypeWithX(v interface{})  {
	t := reflect.TypeOf(v)
	switch t {
	case reflect.TypeOf(float32(1)), reflect.TypeOf(float64(1)):
		fmt.Println("Float")
	case reflect.TypeOf(int16(1)), reflect.TypeOf(int64(1)), reflect.TypeOf(int8(1)), reflect.TypeOf(int32(1)), reflect.TypeOf(int(1)):
		fmt.Println("Integer")
	case reflect.TypeOf(MyCat{}):
		fmt.Println("struct X")
	default:
		fmt.Println("Unknown", t)
	}
}

type MyCat struct {
	Name string	`json:"name" kk:"lovely"`
	Age int
}

func (c *MyCat)UpdateAge(i int) {
	c.Age = i
}

func (c *MyCat)Update(i int, n string) {
	c.Age = i
	c.Name = n
}

func main() {
	a := 2
	b := MyCat{Name: "kk", Age: 5}
	checkType(a)
	checkType(b)

	checkTypeWithX(a)
	checkTypeWithX(b)

	fmt.Println("MyCat Age is ", reflect.ValueOf(b).FieldByName("Age"))
	reflect.ValueOf(&b).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	fmt.Println("kk ", b)

	reflect.ValueOf(&b).MethodByName("Update").Call([]reflect.Value{reflect.ValueOf(2), reflect.ValueOf("xixi")})
	fmt.Println("kk ", b)

	nameFiled, ok := reflect.TypeOf(b).FieldByName("Name")
	if ok {
		fmt.Println("json value = ", nameFiled.Tag.Get("json"))
		fmt.Println("kk value = ", nameFiled.Tag.Get("kk"))
	}else {
		fmt.Println("Failed to get Name field")
	}
}
