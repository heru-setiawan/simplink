package repository

import (
	"context"
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
