package sql

import (
	"github.com/ribeirosaimon/aergia-utils/domain/entities/role"
	"github.com/ribeirosaimon/aergia-utils/domain/vo"
)

type User struct {
	ID          vo.UUID         `json:"id"`
	Username    vo.Name         `json:"username"`
	Password    vo.Password     `json:"password"`
	Email       vo.Email        `json:"email"`
	FirstName   vo.Name         `json:"first_name"`
	LastName    vo.Name         `json:"last_name"`
	Status      vo.Status       `json:"status"`
	Role        role.AergiaRole `json:"role"`
	LoginAtempt uint8           `json:"login_attempt"`
	Audit
}

func NewUser(userName, password, email, firstName, lastName string) (*User, error) {
	user := User{}

	newPassword, err := vo.NewPassword(password)
	if err != nil {
		return nil, err
	}
	if err = newPassword.Encode(); err != nil {
		return nil, err
	}
	user.Password = *newPassword

	newEmail, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}
	user.Email = *newEmail

	user.ID = *vo.NewUuid()
	user.Username = *vo.NewName(userName)
	user.FirstName = *vo.NewName(firstName)
	user.LastName = *vo.NewName(lastName)

	user.Role = role.USER
	user.Status = vo.PENDING
	user.LoginAtempt = 0

	return &user, nil
}
