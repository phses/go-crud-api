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
	ID          int
	Title       string
	Genre       GenreType
	Author      string
	ReleaseDate time.Time
}
