package valueobject

type Status string

const (
	ACTIVE   Status = "Active"
	INACTIVE Status = "Inactive"
	PENDING  Status = "Pending"
)
