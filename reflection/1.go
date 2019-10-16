package reflection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var jsonString = `
{
 "a":"b",
 "c":1,
 "d":[1,2,"b"],
 "e": {
   "f":"g"
 }
}
`

func f1() {
	var unknown map[string]interface{}
	fmt.Println(json.Unmarshal([]byte(jsonString), &unknown))

	for key, value := range unknown {
		fmt.Println(key, ":", reflect.TypeOf(value))
	}
}