package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// FillBySettings 以map中的同名字段填充结构体
func FillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("第一个参数必须为指针类型")
	}
	// Elem() 获取指针指向的值
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("第一个参数必须为结构体的指针类型")
	}

	var field reflect.StructField
	var ok bool

	for k, v := range settings {
		if field, ok = reflect.ValueOf(st).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)	// 这是指针
			vstr = vstr.Elem()	// 这是结构体
			vstr.FieldByName(k).Set(reflect.ValueOf(v))	// 给结构体赋值
		}
	}

	return nil
}

func TestFillBySettings(t *testing.T)  {
	var a MyCat
	err := FillBySettings(&a, map[string]interface{}{
		"Name": "kk",
	})
	if err != nil {
		fmt.Println("err = ", err)
	}

	fmt.Println("after fill a = ", a)
}