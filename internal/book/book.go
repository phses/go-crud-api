package book

import (
	"time"
)

type GenreType string

const (
	Terror         GenreType = "terror"
	Romance        GenreType = "romance"
	ScienceFiction GenreType = "science-fiction"
	Novel          GenreType = "novel"
)

type Book struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Genre       GenreType `db:"genre"`
	Author      string    `db:"author"`
	ReleaseDate time.Time `db:"release_date"`
}
