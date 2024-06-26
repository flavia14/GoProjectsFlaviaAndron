package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"Bookstore/bookstorePack"
	"testing"
	"log"
	"io"
	"bytes"
	"os"
)


func TestDatabaseConnect(t *testing.T) {

	bookstorePack.DatabaseConnect()

	databaseTest, err := sql.Open("sqlite3", "./bookTable.db")

	defer databaseTest.Close()

	rows, err := databaseTest.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='books'")
	if err != nil {
		t.Fatalf("Failed to query the database: %v", err)
	}
	
	defer rows.Close()

	if !rows.Next() {
		t.Error("Expected the 'books' table to exist, but it does not")
	}
}

func TestCreateBook(t *testing.T) {
	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	firstBook := bookstorePack.Book{
		Title: "g",
		Author: "Ion ",
		Genre: "drama",
		Price: 12,
	}

	_, err = database.Exec(`CREATE TABLE books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		genre TEXT,
		price INTEGER
	)`)
	if err != nil {
		t.Fatal(err)
	}

	bookstorePack.CreateBook(firstBook)

	rows, err := database.Query("SELECT COUNT(*) FROM books")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			t.Fatal(err)
		}
	}

	if count != 0 {
		t.Errorf("Expected 1 book, got %d", count)
	}
}
