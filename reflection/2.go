package reflection

import (
	"fmt"
	"reflect"
)

// Type & Kind

func f2() {

	var a MyStruct

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())
}
