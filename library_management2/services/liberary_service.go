package services

import (
	"errors"
	"library_management2/models"
	"sync"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ReserveBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type ReservationRequest struct {
	BookID   int
	MemberID int
	Result   chan error
}


type Library struct {
	Books       map[int]models.Book
	Members     map[int]models.Member
	mutex       sync.Mutex
	Reservation chan ReservationRequest
}

func NewLibrary() *Library {
	lib := &Library{
		Books:       make(map[int]models.Book),
		Members:     make(map[int]models.Member),
		Reservation: make(chan ReservationRequest, 100),
	}

	// Start the reservation worker
	go lib.ReservationWorker()
	return lib
}


func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}


func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}


func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	return nil
}



func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	book.Status = "Available"
	l.Books[bookID] = book
	l.Members[memberID] = member

	return nil
}



func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}


func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}


func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = member
}


