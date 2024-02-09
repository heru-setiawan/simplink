package service

import (
	"context"
	"errors"
	"simplink/features/links/core"
	"simplink/features/links/mocks"
	"simplink/helpers/exceptions"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLink(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewLinkService(repo)
	ctx := context.Background()

	t.Run("empty destination link", func(t *testing.T) {
		caseData := core.Link{}

		result, err := srv.Create(ctx, caseData)

		assert.ErrorAs(t, err, &exceptions.Validation{})
		assert.ErrorContains(t, err, "destination")
		assert.Nil(t, result)
	})

	t.Run("to long short link", func(t *testing.T) {
		caseData := core.Link{
			Short:       "exampleOfCustomShortLink",
			Destination: "https://www.google.com",
		}

		result, err := srv.Create(ctx, caseData)

		assert.ErrorAs(t, err, &exceptions.Validation{})
		assert.ErrorContains(t, err, "short")
		assert.Nil(t, result)
	})

	t.Run("short already exist", func(t *testing.T) {
		caseData := core.Link{
			Short:       "example",
			Destination: "https://www.google.com",
			ExpiredAt:   time.Now().Add(time.Duration(time.Now().Year() + 5)),
		}

		repo.On("Create", ctx, caseData).Return(exceptions.NewRepository("custom short link already exist")).Once()

		result, err := srv.Create(ctx, caseData)

		assert.ErrorAs(t, err, &exceptions.Repository{})
		assert.ErrorContains(t, err, "short")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("errors from repository", func(t *testing.T) {
		caseData := core.Link{
			Short:       "example",
			Destination: "https://www.google.com",
			ExpiredAt:   time.Now().Add(time.Duration(time.Now().Year() + 5)),
		}

		repo.On("Create", ctx, caseData).Return(errors.New("some error from repository")).Once()

		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success with duplicated random short", func(t *testing.T) {
		caseData := core.Link{
			Destination: "https://www.google.com",
			ExpiredAt:   time.Now().Add(time.Duration(time.Now().Year() + 5)),
		}

		repoCaseArg := mock.MatchedBy(func(i interface{}) bool {
			link := i.(core.Link)
			return link.Destination == caseData.Destination
		})
		repo.On("Create", ctx, repoCaseArg).Return(exceptions.NewRepository("custom short link already exist")).Once()
		repo.On("Create", ctx, repoCaseArg).Return(nil).Once()

		result, err := srv.Create(ctx, caseData)

		assert.NoError(t, err)
		assert.Equal(t, caseData.Destination, result.Destination)

		repo.AssertExpectations(t)
	})

	t.Run("success with custom short", func(t *testing.T) {
		caseData := core.Link{
			Short:       "example",
			Destination: "https://www.google.com",
			ExpiredAt:   time.Now().Add(time.Duration(time.Now().Year() + 5)),
		}

		repo.On("Create", ctx, caseData).Return(nil).Once()

		result, err := srv.Create(ctx, caseData)

		assert.NoError(t, err)
		assert.Equal(t, &caseData, result)

		repo.AssertExpectations(t)
	})
}
