package book

type Book struct {
	ID   int
	Name *string
}

var (
	books []*Book = []*Book{}
	ID            = 1
)

func GetBooks() []*Book {
	return books
}
