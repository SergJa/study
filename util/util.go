package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func FailOnError(err error, message string) {
	if err == nil {
		return
	}
	fmt.Println(message)
	panic(fmt.Sprintf("%s: %s", message, err))
}

// красивый вывод в json-style
func PrettyPrint(data interface{}) {
	dataJson, _ := json.Marshal(data)

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, dataJson, "", "\t")

	fmt.Println(prettyJSON.String())

}