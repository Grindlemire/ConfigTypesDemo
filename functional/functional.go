package main

import (
	"fmt"

	"github.schq.secious.com/UltraViolet/configPatterns/sub"
)

// ConfigFunc is the function that will configure the struct opts
type ConfigFunc func(*MyType) error

// MyType ...
type MyType struct {
	MyStr string
	MyInt int
	MySub *sub.Thing

	internalInt int
}

// Speak ...
func (t MyType) Speak() {
	fmt.Printf("\n%s\n", t.MyStr)
	fmt.Printf("My Int: %d\n", t.MyInt)
	fmt.Printf("Sub is: %#v\n\n", t.MySub)
}

func main() {
	functionalType1, _ := NewMyTypeFunctional(
		SetMyStr("Functional Constructor passing in object"),
		SetMyInt(1339),
		SetMySub(sub.NewThing(12, "My sub string functional")),
	)

	defaultFunctionalType, _ := NewMyTypeFunctional()

	functionalType2, _ := NewMyTypeFunctional(
		SetMyStr("Functional Constructor passing in args"),
		SetMyInt(1339),
		SetMySubAlternative("My sub string functional alternative", 12),
	)

	functionalType1.Speak()
	functionalType2.Speak()
	defaultFunctionalType.Speak()
}

// NewMyTypeFunctional ...
func NewMyTypeFunctional(opts ...ConfigFunc) (t *MyType, err error) {
	t = &MyType{
		MyStr: "Sane functional default",
		MyInt: 2,
		MySub: sub.NewThing(1, "Sane functional substring default"),
	}

	for _, o := range opts {
		err = o(t)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

// SetMyStr returns a function that satisfies the ConfigFunc type
func SetMyStr(str string) ConfigFunc {
	return func(t *MyType) error {
		t.MyStr = str
		return nil
	}
}

// SetMyInt returns a function that satisfies the ConfigFunc type
func SetMyInt(i int) ConfigFunc {
	return func(t *MyType) error {
		t.MyInt = i
		return nil
	}
}

// SetMySub returns a function that satisfies the ConfigFunc type
func SetMySub(s *sub.Thing) ConfigFunc {
	return func(t *MyType) error {
		t.MySub = s
		return nil
	}
}

// SetMySubAlternative returns a function that satisfies the ConfigFunc type
func SetMySubAlternative(str string, i int) ConfigFunc {
	return func(t *MyType) error {
		s := sub.NewThing(i, str)
		// do other stuff here to initialize
		t.MySub = s
		return nil
	}
}
