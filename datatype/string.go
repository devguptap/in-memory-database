package datatype

import (
	"errors"
	"fmt"
)

type String struct{}

const maxStringLength = 20

func (s *String) IsValidDataType(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func (s *String) ValidateValue(value interface{}) error {
	if !s.IsValidDataType(value) {
		return errors.New("invalid data type. data should be of string type")
	}

	sValue := value.(string)

	if !s.isValidLength(sValue) {
		return errors.New(fmt.Sprintf("Invalid length : %d. Max Length supported : %d", len(value.(string))))
	}
	return nil
}

func (s *String) isValidLength(value string) bool {
	return len(value) <= maxStringLength
}
