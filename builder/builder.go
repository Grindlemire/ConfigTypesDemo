package main

import (
	"fmt"

	"github.schq.secious.com/UltraViolet/configPatterns/sub"
)

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
	builderType1 := NewMyType().
		SetMyStr("Builder Constructor passing in obj").
		SetMyInt(1338).
		SetMySub(sub.NewThing(11, "My sub string builder"))

	builderType2 := NewMyType().
		SetMyStr("Builder Constructor passing in args").
		SetMyInt(1338).
		SetMySub(sub.NewThing(11, "My sub string builder"))

	defaultBuilderType := NewMyType()

	builderType1.Speak()
	builderType2.Speak()
	defaultBuilderType.Speak()
}

// NewMyType ...
func NewMyType() MyType {
	return MyType{
		MyStr: "Sane builder Default",
		MyInt: 1,
		MySub: &sub.Thing{
			SubStr: "Sane builder substring default",
			SubInt: 1,
		},
	}
}

// SetMyStr ...
func (t MyType) SetMyStr(str string) MyType {
	t.MyStr = str
	return t
}

// SetMyInt ...
func (t MyType) SetMyInt(i int) MyType {
	t.MyInt = i
	return t
}

// SetMySub ...
func (t MyType) SetMySub(sub *sub.Thing) MyType {
	t.MySub = sub
	return t
}

// SetMySubAlternative ...
func (t MyType) SetMySubAlternative(str string, i int) MyType {
	sub := sub.NewThing(i, str)
	t.MySub = sub
	return t
}
