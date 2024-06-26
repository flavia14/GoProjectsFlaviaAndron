package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"Bookstore/bookstorePack"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Title  string
	Author string
	Genre  string
	Price  string
}

func main() {
	createFlag := flag.Bool("create", false, "Create a book")
	readFlag := flag.Bool("read", false, "Read books")
	updateFlag := flag.Bool("update", false, "Update a book")
	deleteFlag := flag.Bool("delete", false, "Delete a book")
	idFlag := flag.Int("id", 0, "ID of the book to update/delete")

	flag.Parse()

	database, err := sql.Open("sqlite3", "./bookTable.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	switch {
	case *createFlag:
		book := Book{
			Title:  flag.Arg(0),
			Author: flag.Arg(1),
			Genre:  flag.Arg(2),
			Price:  flag.Arg(3),
		}
		bookstorePack.CreateBook(database, bookstorePack.Book(book))

	case *readFlag:
		books := bookstorePack.ReadBooks(database)
		for _, book := range books {
			fmt.Println(book)
		}

	case *updateFlag:
		updatedBook := Book{
			Title:  flag.Arg(1),
			Author: flag.Arg(2),
			Genre:  flag.Arg(3),
			Price:  flag.Arg(4),
		}
		bookstorePack.UpdateBook(database, *idFlag, bookstorePack.Book(updatedBook))

	case *deleteFlag:
		bookstorePack.DeleteBook(database, *idFlag)

	default:
		fmt.Println("No operation specified. Please provide a valid flag.")
	}
}
