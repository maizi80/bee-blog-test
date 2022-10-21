package util

import (
	"fmt"
	"reflect"
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

// convert any value to int
func ToInt(v any) (d int, err error) {
	val := reflect.ValueOf(v)
	switch v.(type) {
	case int, int8, int16, int32, int64:
		d = int(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		d = int(val.Uint())
	case float32, float64:
		d = int(val.Float())
	case string:
		d, _ = strconv.Atoi(val.String())
	default:
		err = fmt.Errorf("unknown type `%T`", v)
	}
	return
}

func ToUInt(v any) (d uint, err error) {
	intNum, e := ToInt(v)
	d = uint(intNum)
	err = e
	return
}
