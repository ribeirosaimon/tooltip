package valueobject

import (
	"regexp"

	"github.com/pkg/errors"
)

type Email struct {
	value string
}

func NewEmail(e string) (*Email, error) {
	if isValid(e) {
		return &Email{value: e}, nil
	}
	return nil, errors.New("Invalid email")
}

func isValid(e string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).
		MatchString(e)
}
