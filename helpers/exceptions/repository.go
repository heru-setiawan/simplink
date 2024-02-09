package exceptions

func NewRepository(message string) error {
	return Repository{
		message: message,
	}
}

type Repository struct {
	message string
}

func (err Repository) Error() string {
	return err.message
}
