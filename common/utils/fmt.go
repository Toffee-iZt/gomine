package utils

import (
	"fmt"
	"strings"
)

// ToString converts anything to string
// but unlike fmt.Sprint doesn't put spaces
func ToString(v ...interface{}) string {
	strs := make([]string, len(v))
	for i, str := range v {
		strs[i] = fmt.Sprintf("%v", str)
	}
	return strings.Join(strs, "")
}
