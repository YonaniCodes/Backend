package controllers

import (
	"bufio"
	"fmt"
	"library_management2/models"
	"library_management2/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	Library *services.Library
}

func NewLibraryController(library *services.Library) *LibraryController {
	return &LibraryController{Library: library}
}

func (c *LibraryController) Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Library Management System ===")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			c.AddBook(reader)
		case 2:
			c.RemoveBook(reader)
		case 3:
			c.BorrowBook(reader)
		case 4:
			c.ReturnBook(reader)
		case 5:
			c.ListAvailableBooks()
		case 6:
			c.ListBorrowedBooks(reader)
		case 7:
			c.AddMember(reader)
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func (c *LibraryController) AddBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	id := c.readInt(reader)

	fmt.Print("Enter Book Title: ")
	title := c.readString(reader)

	fmt.Print("Enter Book Author: ")
	author := c.readString(reader)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	c.Library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (c *LibraryController) RemoveBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	id := c.readInt(reader)

	c.Library.RemoveBook(id)
	fmt.Println("Book removed!")
}

func (c *LibraryController) BorrowBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	bookID := c.readInt(reader)

	fmt.Print("Enter Member ID: ")
	memberID := c.readInt(reader)

	err := c.Library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (c *LibraryController) ReturnBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	bookID := c.readInt(reader)

	fmt.Print("Enter Member ID: ")
	memberID := c.readInt(reader)

	err := c.Library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned!")
	}
}

func (c *LibraryController) ListAvailableBooks() {
	fmt.Println("\nAvailable Books:")
	books := c.Library.ListAvailableBooks()
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func (c *LibraryController) ListBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	memberID := c.readInt(reader)

	books := c.Library.ListBorrowedBooks(memberID)
	fmt.Println("\nBorrowed Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func (c *LibraryController) AddMember(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	id := c.readInt(reader)

	fmt.Print("Enter Member Name: ")
	name := c.readString(reader)

	member := models.Member{
		ID:            id,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}

	c.Library.AddMember(member)
	fmt.Println("Member added successfully!")
}

// Helper functions
func (c *LibraryController) readString(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

func (c *LibraryController) readInt(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	value, _ := strconv.Atoi(strings.TrimSpace(input))
	return value
}
