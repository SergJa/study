package recover

import (
	"errors"
	"fmt"
)

// правильная обработка ошибок

type MyError2 error

func Isterika2() error {
	err := MyError(errors.New("Осень!"))
	if  err != nil {
		return err
	}
	//...
	return nil
}

func f2() {
	err := Isterika2()
	if err != nil {
		fmt.Println("Найдена ошибка:", err.Error())
	}
}