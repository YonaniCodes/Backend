package models

// Member represents a library member with a collection of borrowed books.
type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}
