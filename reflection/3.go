package reflection

// reflection структуры

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	A string `json:"abc"`
	B int
}

func (m MyStruct) MyMethod(a string, b int) {
	fmt.Println(m.A, a, b + m.B)
}

func f3() {
	var unknown interface{}//MyStruct{}

	var unknownType = reflect.TypeOf(unknown)

	if unknownType.Kind() != reflect.Struct {
		fmt.Println("Structure required")
		return
	}

	// работа с тэгами
	field, exists := unknownType.FieldByName("A")
	fmt.Println(exists)
	fmt.Println(field.Tag.Lookup("json"))

	// получение имен всех полей
	fmt.Println("Поля MyStruct")
	for i:=0; i<unknownType.NumField(); i++ {
		field := unknownType.Field(i)
		fmt.Println(field.Name, " - ", field.Type.Name())
	}

	fmt.Println("Методы MyStruct")
	for i:=0; i<unknownType.NumMethod(); i++ {
		method := unknownType.Method(i)
		fmt.Println(method.Name)
	}
}
