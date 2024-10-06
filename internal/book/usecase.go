package book

import (
	"context"
	"errors"
	"time"
)

type UseCase interface {
	Get(ctx context.Context, id int) (*Book, error)
	Create(ctx context.Context, title string, genre string, author string, releaseDate time.Time) (int, error)
}

type BookUseCase struct {
	repo Repository
}

func NewBookUseCase(r Repository) UseCase {
	return &BookUseCase{repo: r}
}

func (uc *BookUseCase) Get(ctx context.Context, id int) (*Book, error) {
	if id <= 0 {
		return nil, errors.New("invalid Id")
	}

	return uc.repo.GetById(ctx, id)
}

func (uc *BookUseCase) Create(ctx context.Context, title string, genre string, author string, releaseDate time.Time) (int, error) {
	if title == "" || genre == "" || author == "" {
		return 0, errors.New("missing required fields")
	}

	book := &Book{
		Author:      author,
		Title:       title,
		Genre:       GenreType(genre),
		ReleaseDate: releaseDate,
	}

	return uc.repo.Create(ctx, book)
}
