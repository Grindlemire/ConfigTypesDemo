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

// MyTypeConf used for the config style
type MyTypeConf struct {
	MyStr string
	MyInt int
	MySub *sub.Thing
}

func main() {
	myTypeConf := MyTypeConf{
		MyStr: "Config constructor",
		MyInt: 1337,
		MySub: sub.NewThing(10, "My sub string config"),
	}
	configType := NewMyType(myTypeConf)

	configType.Speak()

}

// NewMyType ...
func NewMyType(conf MyTypeConf) MyType {
	return MyType{
		MyStr: conf.MyStr,
		MyInt: conf.MyInt,
		MySub: conf.MySub,

		internalInt: 99,
	}
}
