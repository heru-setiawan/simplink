package service

import (
	"context"
	"errors"
	"simplink/features/links/core"
	"simplink/helpers/exceptions"
	"time"
)

func NewLinkService(repository core.Repository) core.Service {
	return &linkService{
		repository: repository,
	}
}

type linkService struct {
	repository core.Repository
}

func (srv *linkService) Create(ctx context.Context, data core.Link) (*core.Link, error) {
	var generatedShort bool

	data.SetExpired(time.Duration(time.Now().Year() + 5))

	if data.Short == "" {
		data.GenerateShort(6)
		generatedShort = true
	}

	if err := validateCreateLink(data); err != nil {
		return nil, err
	}

	if err := srv.repository.Create(ctx, data); err != nil {
		if errors.As(err, &exceptions.Repository{}) && generatedShort {
			data.Short = ""
			return srv.Create(ctx, data)
		}

		return nil, err
	}

	return &data, nil
}
