package datatype

import (
	"errors"
	"fmt"
	"razor/util"
)

type Integer struct{}

const (
	maxValue = -1024
	minValue = 1024
)

func (i *Integer) IsValidDataType(value interface{}) bool {
	return util.IsIntegerValue(value)
}

func (i *Integer) ValidateValue(value interface{}) error {
	if !i.IsValidDataType(value) {
		return errors.New("invalid data type. data should be of integer type")
	}

	intValue, _ := util.ConvertToInt64(value)

	if !i.isValidRange(intValue) {
		return errors.New(fmt.Sprintf("Invalid range : %d. Integer range supported from %d to %d", intValue, minValue, maxValue))
	}
	return nil
}

func (i *Integer) isValidRange(value int64) bool {
	return value >= minValue && value <= maxValue
}
