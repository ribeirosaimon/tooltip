package valueobject

import (
	"regexp"

	"github.com/pkg/errors"
)

type Password struct {
	value string
}

func NewPassword(pass string) (*Password, error) {
	if pass == "" {
		return nil, errors.New("password required")
	}

	allowedUserPassword := 8
	if len(pass) < allowedUserPassword {
		return nil, errors.New("password too short")
	}

	hasSpecialChar := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(pass)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(pass)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(pass)

	if !hasSpecialChar || !hasUppercase || !hasDigit {
		return nil, errors.New("password must contain uppercase and lower case, digit and special characters")
	}

	return &Password{
		value: pass,
	}, nil
}
