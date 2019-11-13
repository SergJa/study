package reflection

// рефлексия Value
import "reflect"

func F4() {
	var unknown = &MyStruct{"A in struct", 7}

	var unknownValue = reflect.ValueOf(unknown)

	// вызов метода
	unknownValue.Method(0).Call([]reflect.Value{
		reflect.ValueOf("a"),
		reflect.ValueOf(2),
	})

	// получение содержимого указателя
	unknownValue.Elem()

	// получение базового значения (только если уверены в типе)
	unknownValue.Bool()

	// заменить значение
	unknownValue.Set(reflect.ValueOf(&MyStruct{A:"NewA"}))
}
