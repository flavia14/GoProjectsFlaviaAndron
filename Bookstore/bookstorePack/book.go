package bookstorePack

import(
	"errors"
)

func ValidateBookTitle(title string) error {
	if title == "" {
		return errors.New("book title cannot be empty")
	}
	return nil
}

func ValidateBookAuthor(author string) error {
	if author == "" {
		return errors.New("book author cannot be empty")
	}
	return nil
}

func ValidateBookPrice(price int) error {
	if price <= 0 {
		return errors.New("book price must be greater than zero")
	}
	return nil
}