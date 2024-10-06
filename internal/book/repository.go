package book

import "context"

type Repository interface {
	GetById(ctx context.Context, id int) (*Book, error)
	Create(ctx context.Context, book *Book) (int, error)
}
