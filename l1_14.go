package main

import (
	"fmt"
	"reflect"
)

type A struct{}

func main() {
	var i_x interface{} = 3
	var str_x interface{} = "hello"
	var f_x interface{} = 52.52
	var bool_x interface{} = false
	var chan_x interface{} = make(chan *A)

	xs := []interface{}{i_x, str_x, f_x, bool_x, chan_x}

	for _, x := range xs {
		// Узнать тип данных переменной можно с помощью библиотеки reflect и метода reflect.TypeOf
		fmt.Printf("x = %v, type of x = %v\n", x, reflect.TypeOf(x))
	}
}
