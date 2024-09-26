package valueobject

type Name string

func NewName(name string) Name {
	return Name(name)
}
