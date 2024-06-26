package bookstorePack

import (
	"database/sql"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Genre  string
	Price  string
}

func CreateBook(database *sql.DB, book Book) {
	statement, err := database.Prepare("INSERT INTO books (title, author, genre, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Failed to prepare statement:", err)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(book.Title, book.Author, book.Genre, book.Price)
	if err != nil {
		fmt.Println("Failed to create book:", err)
		return
	}

	fmt.Println("Book created successfully!")
}

func ReadBooks(database *sql.DB) []Book {
	rows, err := database.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println("Failed to fetch books:", err)
		return nil
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Title, &book.Author, &book.Genre, &book.Price)
		if err != nil {
			fmt.Println("Failed to scan book:", err)
			continue
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Failed to iterate over books:", err)
		return nil
	}

	return books
}

func UpdateBook(database *sql.DB, id int, book Book) {
	statement, err := database.Prepare("UPDATE books SET title=?, author=?, genre=?, price=? WHERE id=?")
	if err != nil {
		fmt.Println("Failed to prepare statement:", err)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(book.Title, book.Author, book.Genre, book.Price, id)
	if err != nil {
		fmt.Println("Failed to update book:", err)
		return
	}

	fmt.Println("Book updated successfully!")
}

func DeleteBook(database *sql.DB, id int) {
	statement, err := database.Prepare("DELETE FROM books WHERE id=?")
	if err != nil {
		fmt.Println("Failed to prepare statement:", err)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		fmt.Println("Failed to delete book:", err)
		return
	}

	fmt.Println("Book deleted successfully!")
}
