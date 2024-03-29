package core

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(ctx echo.Context) error
	GetByShort(ctx echo.Context) error
}

type Service interface {
	Create(ctx context.Context, data Link) (*Link, error)
	GetByShort(ctx context.Context, short string) (*Link, error)
}

type Repository interface {
	Create(ctx context.Context, data Link) error
	GetByShort(ctx context.Context, short string) (*Link, error)
}
