package utils

import "fmt"

func ErrHandler (err error, data interface{})map[string]interface{} {
	if err != nil {
		return map[string]interface{}{
			"code": 40040,
			"errmsg": fmt.Sprintf("%s", err),
		}
	}else {
		return map[string]interface{}{
			"code": 0,
			"data": data,
		}
	}
}
