package util

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func ContainsIntInString(s string, search int64) bool {
	slice := strings.Split(s, ",")
	for _, v := range slice {
		vInt, _ := strconv.Atoi(v)
		if vInt == int(search) {
			return true
		}
	}
	return false
}
