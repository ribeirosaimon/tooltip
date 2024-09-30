package vo

import (
	"regexp"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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

func (p *Password) Encode() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.value), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.value = string(hash)
	return nil
}

func (p *Password) VerifyPassword(pass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword(hash, []byte(p.value)); err != nil {
		return err
	}
	return nil
}

func (p *Password) GetValue() string {
	return p.value
}
