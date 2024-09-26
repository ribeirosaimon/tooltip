package valueobject

import (
	"github.com/google/uuid"
	"github.com/ribeirosaimon/aergia-utils/logs"
)

type UUID string

func NewUuid() UUID {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		logs.ERROR.Message("Error creating new UUID")
	}
	return UUID(newUUID.String())
}
