package recover

import (
	"errors"
	"fmt"
)

// правильная обработка ошибок

type MyError2 error

func Isterika2() (string, error) {
	err := MyError(errors.New("Осень!"))
	if  err != nil {
		return "", err
	}
	//...
	return "abc", nil
}

func f2() {
	abc, err := Isterika2()
	if err != nil {
		fmt.Println("Найдена ошибка:", err.Error())
		return
	}
	fmt.Println(abc)
}