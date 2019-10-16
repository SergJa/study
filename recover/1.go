package recover

import (
	"errors"
	"fmt"
)

type MyError error

func Isterika() {
	err := MyError(errors.New("Осень!"))
	panic(err)
}

func f() {
	defer func() {
		rec := recover()
		_, isMyError := rec.(MyError)
		if isMyError {
			fmt.Println(rec)
			// логика восстановления
		}
	}()
	Isterika()
}

func main() {
	f()
}

