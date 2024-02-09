package service

import (
	"simplink/features/links/core"
	"simplink/helpers/exceptions"
)

func validateGetByShortLink(short string) error {
	if short == "" {
		return exceptions.NewValidation("short link must be filled in")
	}

	return nil
}

func validateCreateLink(data core.Link) error {
	if data.Destination == "" {
		return exceptions.NewValidation("destination link must be filled in")
	}

	if len(data.Short) > 16 {
		return exceptions.NewValidation("custom short link must less than 16 character")
	}

	return nil
}
