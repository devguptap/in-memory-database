package validation

type IValidation interface {
	Validate(value interface{}, args ...interface{}) bool
	GetErrorMessage() string
}
