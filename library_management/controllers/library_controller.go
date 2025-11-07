package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

type memberRegistry interface {
	AddMember(member models.Member)
	Members() []models.Member
}

// LibraryController coordinates console interactions with the library service layer.
type LibraryController struct {
	manager  services.LibraryManager
	registry memberRegistry
	reader   *bufio.Reader
}

// NewLibraryController creates a controller for the supplied manager implementation.
func NewLibraryController(manager services.LibraryManager) *LibraryController {
	registry, _ := manager.(memberRegistry)
	return &LibraryController{
		manager:  manager,
		registry: registry,
		reader:   bufio.NewReader(os.Stdin),
	}
}

// Run starts the interactive console loop.
func (c *LibraryController) Run() {
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		if c.registry != nil {
			fmt.Println("7. Register Member")
			fmt.Println("8. List Members")
			fmt.Println("9. Exit")
		} else {
			fmt.Println("7. Exit")
		}
		fmt.Print("Choose an option: ")

		choice, err := c.readInt()
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		if c.registry != nil {
			switch choice {
			case 1:
				c.handleAddBook()
			case 2:
				c.handleRemoveBook()
			case 3:
				c.handleBorrowBook()
			case 4:
				c.handleReturnBook()
			case 5:
				c.handleListAvailable()
			case 6:
				c.handleListBorrowed()
			case 7:
				c.handleRegisterMember()
			case 8:
				c.handleListMembers()
			case 9:
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("Unknown option. Please select a valid menu number.")
			}
		} else {
			switch choice {
			case 1:
				c.handleAddBook()
			case 2:
				c.handleRemoveBook()
			case 3:
				c.handleBorrowBook()
			case 4:
				c.handleReturnBook()
			case 5:
				c.handleListAvailable()
			case 6:
				c.handleListBorrowed()
			case 7:
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("Unknown option. Please select a valid menu number.")
			}
		}
	}
}

func (c *LibraryController) handleAddBook() {
	id, err := c.promptInt("Enter book ID: ")
	if err != nil {
		fmt.Println("Invalid ID. Book not added.")
		return
	}
	title := c.promptString("Enter book title: ")
	author := c.promptString("Enter book author: ")

	c.manager.AddBook(models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	})
	fmt.Println("Book added successfully.")
}

func (c *LibraryController) handleRemoveBook() {
	id, err := c.promptInt("Enter book ID to remove: ")
	if err != nil {
		fmt.Println("Invalid ID. No book removed.")
		return
	}
	c.manager.RemoveBook(id)
	fmt.Println("Book removed if it existed in the catalogue.")
}

func (c *LibraryController) handleBorrowBook() {
	bookID, err := c.promptInt("Enter book ID to borrow: ")
	if err != nil {
		fmt.Println("Invalid book ID.")
		return
	}
	memberID, err := c.promptInt("Enter member ID: ")
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}
	if err := c.manager.BorrowBook(bookID, memberID); err != nil {
		fmt.Printf("Unable to borrow book: %v\n", err)
		return
	}
	fmt.Println("Book borrowed successfully.")
}

func (c *LibraryController) handleReturnBook() {
	bookID, err := c.promptInt("Enter book ID to return: ")
	if err != nil {
		fmt.Println("Invalid book ID.")
		return
	}
	memberID, err := c.promptInt("Enter member ID: ")
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}
	if err := c.manager.ReturnBook(bookID, memberID); err != nil {
		fmt.Printf("Unable to return book: %v\n", err)
		return
	}
	fmt.Println("Book returned successfully.")
}

func (c *LibraryController) handleListAvailable() {
	books := c.manager.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books are currently available.")
		return
	}
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) handleListBorrowed() {
	memberID, err := c.promptInt("Enter member ID: ")
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}
	books := c.manager.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books found for this member or member does not exist.")
		return
	}
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) handleRegisterMember() {
	if c.registry == nil {
		fmt.Println("Member registration is not supported in this configuration.")
		return
	}
	id, err := c.promptInt("Enter member ID: ")
	if err != nil {
		fmt.Println("Invalid ID. Member not added.")
		return
	}
	name := c.promptString("Enter member name: ")
	c.registry.AddMember(models.Member{ID: id, Name: name})
	fmt.Println("Member registered successfully.")
}

func (c *LibraryController) handleListMembers() {
	if c.registry == nil {
		fmt.Println("Member listing is not supported in this configuration.")
		return
	}
	members := c.registry.Members()
	if len(members) == 0 {
		fmt.Println("No members registered.")
		return
	}
	fmt.Println("Registered Members:")
	for _, member := range members {
		fmt.Printf("ID: %d | Name: %s | Borrowed Books: %d\n", member.ID, member.Name, len(member.BorrowedBooks))
	}
}

func (c *LibraryController) promptInt(prompt string) (int, error) {
	fmt.Print(prompt)
	input, err := c.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (c *LibraryController) promptString(prompt string) string {
	fmt.Print(prompt)
	input, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (c *LibraryController) readInt() (int, error) {
	input, err := c.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}
	return value, nil
}
