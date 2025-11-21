package models

type BookStatus string

// Step 2: Define constants for BookStatus
const (
	Borrowed  BookStatus = "Borrowed"
	Available BookStatus = "Available"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Status   BookStatus
	Reserved bool // true if reserved
}
