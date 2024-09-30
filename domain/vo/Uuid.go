package vo

import (
	"github.com/google/uuid"
	"github.com/ribeirosaimon/aergia-utils/logs"
)

type UUID struct {
	value string
}

func NewUuid() *UUID {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		logs.ERROR.Message("Error creating new UUID")
	}
	return &UUID{
		value: newUUID.String(),
	}
}

func (uuid *UUID) GetUUID() string {
	return uuid.value
}
