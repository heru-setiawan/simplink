package exceptions

func NewValidation(message string) error {
	return Validation{
		message: message,
	}
}

type Validation struct {
	message string
}

func (err Validation) Error() string {
	return err.message
}
