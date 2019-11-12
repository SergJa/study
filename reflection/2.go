package reflection

import (
	"fmt"
	"reflect"
)

// Type & Kind

type MyInt int

func F2() {
	var a MyStruct

	var i MyInt
	var j int

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())

	fmt.Println(reflect.TypeOf(i).Kind() == reflect.TypeOf(j).Kind())
}
