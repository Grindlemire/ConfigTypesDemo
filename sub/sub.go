package sub

// Thing ...
type Thing struct {
	SubStr string
	SubInt int
}

// NewThing ...
func NewThing(subInt int, subStr string) *Thing {
	return &Thing{
		SubStr: subStr,
		SubInt: subInt,
	}
}
