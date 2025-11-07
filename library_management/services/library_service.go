package services

import (
	"errors"
	"sort"

	"library_management/models"
)

const (
	statusAvailable = "Available"
	statusBorrowed  = "Borrowed"
)

// LibraryManager defines the operations available for managing the library.
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

// Library implements LibraryManager using in-memory data stores.
type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

// NewLibrary creates a new Library instance with optional seed data.
func NewLibrary(books []models.Book, members []models.Member) *Library {
	bookMap := make(map[int]models.Book, len(books))
	for _, book := range books {
		if book.Status == "" {
			book.Status = statusAvailable
		}
		bookMap[book.ID] = book
	}

	memberMap := make(map[int]models.Member, len(members))
	for _, member := range members {
		// Ensure member borrowed book statuses align with central store.
		cleaned := make([]models.Book, 0, len(member.BorrowedBooks))
		for _, book := range member.BorrowedBooks {
			if b, ok := bookMap[book.ID]; ok {
				b.Status = statusBorrowed
				bookMap[b.ID] = b
				cleaned = append(cleaned, b)
			}
		}
		member.BorrowedBooks = cleaned
		memberMap[member.ID] = member
	}

	return &Library{
		books:   bookMap,
		members: memberMap,
	}
}

// AddBook inserts or replaces a book in the catalogue.
func (l *Library) AddBook(book models.Book) {
	if l.books == nil {
		l.books = make(map[int]models.Book)
	}
	if book.Status == "" {
		book.Status = statusAvailable
	}
	l.books[book.ID] = book
}

// RemoveBook deletes a book from the catalogue if it exists.
func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
	for memberID, member := range l.members {
		updated := member.BorrowedBooks[:0]
		for _, borrowed := range member.BorrowedBooks {
			if borrowed.ID != bookID {
				updated = append(updated, borrowed)
			}
		}
		member.BorrowedBooks = updated
		l.members[memberID] = member
	}
}

// AddMember ensures a member exists in the library registry.
func (l *Library) AddMember(member models.Member) {
	if l.members == nil {
		l.members = make(map[int]models.Member)
	}
	l.members[member.ID] = member
}

// BorrowBook allows a member to borrow a book if it is available.
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == statusBorrowed {
		return errors.New("book already borrowed")
	}

	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	book.Status = statusBorrowed
	l.books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}

// ReturnBook allows a member to return a borrowed book.
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}

	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	found := false
	updated := member.BorrowedBooks[:0]
	for _, borrowed := range member.BorrowedBooks {
		if borrowed.ID == bookID {
			found = true
			continue
		}
		updated = append(updated, borrowed)
	}

	if !found {
		return errors.New("book not borrowed by member")
	}

	member.BorrowedBooks = updated
	l.members[memberID] = member

	book.Status = statusAvailable
	l.books[bookID] = book
	return nil
}

// ListAvailableBooks returns the books currently available for borrowing.
func (l *Library) ListAvailableBooks() []models.Book {
	available := make([]models.Book, 0)
	for _, book := range l.books {
		if book.Status == statusAvailable {
			available = append(available, book)
		}
	}
	sort.Slice(available, func(i, j int) bool { return available[i].ID < available[j].ID })
	return available
}

// ListBorrowedBooks returns the books currently borrowed by the specified member.
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.members[memberID]
	if !ok {
		return nil
	}
	borrowed := make([]models.Book, len(member.BorrowedBooks))
	copy(borrowed, member.BorrowedBooks)
	sort.Slice(borrowed, func(i, j int) bool { return borrowed[i].ID < borrowed[j].ID })
	return borrowed
}

// Members returns a snapshot of registered members.
func (l *Library) Members() []models.Member {
	members := make([]models.Member, 0, len(l.members))
	for _, member := range l.members {
		members = append(members, member)
	}
	sort.Slice(members, func(i, j int) bool { return members[i].ID < members[j].ID })
	return members
}
