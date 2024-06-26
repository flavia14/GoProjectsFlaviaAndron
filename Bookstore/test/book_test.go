package test

import (
	_ "github.com/mattn/go-sqlite3"
	"Bookstore/bookstorePack"
	"testing"
)

func TestValidateBookTitle(t *testing.T) {
	err := bookstorePack.ValidateBookTitle("title")
	if err != nil {
		t.Error("Got an error")
	}
}

func TestValidateBookTitleEmpty(t *testing.T) {
	err := bookstorePack.ValidateBookTitle("")
	if err == nil {
		t.Error("Did not get an error")
	}
}

func TestValidateBookAuthor(t *testing.T) {
	err := bookstorePack.ValidateBookAuthor("Author")
	if err != nil {
		t.Error("Got an error")
	}
}

func TestValidateBookAuthorEmpty(t *testing.T) {
	err := bookstorePack.ValidateBookAuthor("")
	if err == nil {
		t.Error("Did not get an error")
	}
}

func TestValidateBookPrice(t *testing.T) {
	err := bookstorePack.ValidateBookPrice(12)
	if err != nil {
		t.Error("Got an error")
	}
}

func TestValidateBookNegativePrice(t *testing.T) {
	err := bookstorePack.ValidateBookPrice(-12)
	if err == nil {
		t.Error("Did not get an error")
	}
}
