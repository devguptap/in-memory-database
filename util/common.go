package util

import "errors"

func ConvertToInt64(value interface{}) (int64, error) {
	var intValue int64
	var err error
	switch value.(type) {
	case int:
		intValue = int64(value.(int))
	case int8:
		intValue = int64(value.(int))
	case int16:
		intValue = int64(value.(int))
	case int32:
		intValue = int64(value.(int))
	case int64:
		intValue = int64(value.(int))
	default:
		err = errors.New("invalid type")
	}
	return intValue, err
}

func IsIntegerValue(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return true
	default:
		return false
	}
}
