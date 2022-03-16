package validator

type Validator interface {
	Validate(value string, values []string) bool
}

type ValidatorImpl struct {}

func (v ValidatorImpl) Validate(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}

