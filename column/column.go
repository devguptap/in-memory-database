package column

import (
	"errors"
	"razor/constants"
	"razor/datatype"
	"razor/validation"
	"strings"
)

type Column struct {
	Name        string
	DataType    datatype.IDataType
	Constraints []validation.IValidation
}

func NewColumn(columnName string, dataType string, constraints []string) Column {
	c := Column{
		Name:        columnName,
		DataType:    getConcreteDataType(dataType),
		Constraints: getConcreteConstraints(constraints),
	}
	return c
}

func getConcreteDataType(dataType string) datatype.IDataType {
	if strings.EqualFold(dataType, constants.IntegerDataType) {
		return &datatype.String{}
	} else {
		return &datatype.Integer{}
	}
}

func getConcreteConstraints(constraints []string) []validation.IValidation {
	var concreteConstraints []validation.IValidation
	for _, constraint := range constraints {
		if strings.EqualFold(constraint, constants.RequiredConstraint) {
			concreteConstraints = append(concreteConstraints, &validation.Required{})
		}
	}
	return concreteConstraints
}

func (c Column) ValidateValueAgainstConstraints(value interface{}) error {
	for _, constraint := range c.Constraints {
		if !constraint.Validate(value) {
			return errors.New(constraint.GetErrorMessage())
		}
	}
	return nil
}
