package variables

import (
	"fmt"
	"reflect"
)

func Variables() {

	a := 3
	fmt.Println("a = ", a)

	var b, c = "Long", 1
	fmt.Println("b = ", b, "c = ", c)

	var d = true
	fmt.Println("d = ", d)

	var e int
	fmt.Println("e = ", e)

	f := "true"
	fmt.Println("f = ", f)

	fmt.Println("========")
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.TypeOf(f))

	//output:
	//a =  3
	//b =  Long c =  1
	//d =  true
	//e =  0
	//f =  true
	//========
	//int
	//string
	//int
	//bool
	//int
	//string

}
