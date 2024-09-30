package vo

type Name struct {
	value string
}

func NewName(name string) *Name {
	return &Name{
		value: name,
	}
}

func (name *Name) GetValue() string {
	return name.value
}
