package reflection

import (
	"fmt"
	"reflect"
)

// работа с тэгами
func F5(struc interface{}) {
	Type := reflect.TypeOf(struc)
	if Type.Kind() == reflect.Struct {
		if field, ok := Type.FieldByName("PersonName"); ok {
			fmt.Println(field.Tag.Lookup("db"))
		}
	}
}