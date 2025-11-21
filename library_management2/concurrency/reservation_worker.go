package services

import (
	"fmt"
	"time"
)

func (l *Library) ReservationWorker() {
	for req := range l.Reservation {
		l.mutex.Lock()
		book, exists := l.Books[req.BookID]
		if !exists {
			req.Result <- fmt.Errorf("book not found")
			l.mutex.Unlock()
			continue
		}

		if book.Status != "Available" || book.Reserved {
			req.Result <- fmt.Errorf("book not available for reservation")
			l.mutex.Unlock()
			continue
		}

		// Mark as reserved
		book.Reserved = true
		l.Books[req.BookID] = book
		req.Result <- nil
		l.mutex.Unlock()

		// Start timer to auto-cancel reservation
		go func(bookID int) {
			time.Sleep(5 * time.Second)
			l.mutex.Lock()
			b := l.Books[bookID]
			if b.Reserved && b.Status == "Available" {
				b.Reserved = false
				l.Books[bookID] = b
				fmt.Printf("Reservation for book %d cancelled (timeout)\n", bookID)
			}
			l.mutex.Unlock()
		}(req.BookID)
	}
}
