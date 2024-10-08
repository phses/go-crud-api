package postgres

import (
	"context"
	"github/phses/go-crud-api/internal/book"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	DB *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) book.Repository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) GetById(ctx context.Context, id int) (*book.Book, error) {
	query := `SELECT id, title, genre, author, release_date FROM books WHERE id = $1`

	var b book.Book

	if err := r.DB.GetContext(ctx, &b, query, id); err != nil {
		return nil, err
	}

	return &b, nil
}

func (r *PostgresRepository) Create(ctx context.Context, book *book.Book) (int, error) {
	query := `INSERT INTO books(title, genre, author, release_date) VALUES($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.DB.QueryRowContext(ctx, query, book.Title, book.Genre, book.Author, book.ReleaseDate).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
