package validation

type Required struct{}

func (r *Required) Validate(value interface{}, args ...interface{}) bool {
	switch value.(type) {
	case int64:
		return value.(int64) != 0
	case string:
		return len(value.(string)) != 0
	default:
		return false
	}
}

func (r *Required) GetErrorMessage() string {
	return "Required validation failed"
}
