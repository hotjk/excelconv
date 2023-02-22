package conv

import (
	"fmt"
	"strings"
	"time"
)

var m map[string]int = map[string]int{}

func Invoke(value string, function string, params []interface{}) (result string) {
	if function == "" {
		result = value
		return
	}
	switch function {
	case "TimeFormat":
		time, _ := time.Parse(params[0].(string), value)
		result = time.Format(params[1].(string))
	case "ToUpper":
		result = strings.ToUpper(value)
	case "Sequence":
		key := params[0].(string)
		value := m[key] + 1
		m[key] = value
		result = fmt.Sprintf(params[1].(string), value)
		//fmt.Println(key, value)
	}
	return
}
