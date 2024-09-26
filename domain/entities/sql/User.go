package sql

import (
	"github.com/ribeirosaimon/aergia-utils/domain/entities/role"
	"github.com/ribeirosaimon/aergia-utils/domain/valueobject"
)

type User struct {
	ID          valueobject.UUID     `json:"id"`
	Username    valueobject.Name     `json:"username"`
	Password    valueobject.Password `json:"password"`
	Email       valueobject.Email    `json:"email"`
	FirstName   valueobject.Name     `json:"first_name"`
	LastName    valueobject.Name     `json:"last_name"`
	Status      valueobject.Status   `json:"status"`
	Role        role.AergiaRole      `json:"role"`
	LoginAtempt uint8                `json:"login_attempt"`
	Audit
}
