package sql

import "time"

type Audit struct {
	UserId    string    `json:"UserId"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
