package book

import (
	"strconv"

	"github.com/gofiber/fiber"
)

// Book structure
type book struct {
	BookId    int    `json:"id"`
	BookName  string `json:"bookname"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

var books []book

// Get all books
func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
	c.JSON(books)
}

// Get a single book
func GetBook(c *fiber.Ctx) {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.Send("Book Id not found")
	}
	for _, book := range books {
		if book.BookId == id {
			c.JSON(book)
		}
	}

}

// Adds a new book
func NewBook(c *fiber.Ctx) {
	var NewBook book
	if err := c.BodyParser(&NewBook); err != nil {
		c.Send("Uncompleted")
	}
	NewBook.BookId = len(books) + 1
	books = append(books, NewBook)
	c.JSON(NewBook)
}

// Update Book
func UpdateBook(c *fiber.Ctx) {
	c.Send("Update a book")
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.Send("Invalid Book Id")
	}
	var updatedBook book
	for i, book := range books {
		if book.BookId == id {
			updatedBook.BookId = id
			books[i] = updatedBook
			c.JSON(updatedBook)
		}
	}
}

// Delete Book
func DeleteBook(c *fiber.Ctx) {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.Send("Invalid Book Id")
	}
	for i, book := range books {
		if book.BookId == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	c.Send("Book deleted")
}
