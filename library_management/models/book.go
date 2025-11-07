package models

// Book represents a single item in the library catalogue.
type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}
