package repository

import (
	"context"
	"errors"
	"simplink/features/links/core"
	"simplink/helpers/exceptions"
	"strings"

	"gorm.io/gorm"
)

func NewLinkRepository(mysql *gorm.DB) core.Repository {
	return &linkRepository{
		mysql: mysql,
	}
}

type linkRepository struct {
	mysql *gorm.DB
}

func (repo *linkRepository) GetByShort(ctx context.Context, short string) (*core.Link, error) {
	mod := new(Link)

	if err := repo.mysql.WithContext(ctx).Where("short = ?", short).First(mod).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewRepository("link not found")
		}

		return nil, err
	}

	return mod.ToEntity(), nil
}

func (repo *linkRepository) Create(ctx context.Context, data core.Link) error {
	mod := newLinkFromEntity(data)

	if err := repo.mysql.WithContext(ctx).Create(&mod).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return exceptions.NewRepository("custom short link already exist")
		}

		return err
	}

	return nil
}
