package book

type Repository interface {
	GetById(id int) (*Book, error)
	Create(book *Book) (int, error)
}
