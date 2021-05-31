package datatype

type IDataType interface {
	IsValidDataType(value interface{}) bool
	ValidateValue(value interface{}) error
}
