package main

import (
	"fmt"

	"github.com/tiago4orion/scherzo/lang"
)

const version = "0.0.1"

func main() {
	fmt.Println("Scherzo compiler ", version)

	value2 := lang.Cons(lang.NewAtom(2), lang.Nil)
	values := lang.Cons(lang.NewAtom(1), lang.NewAtom(value2))

	// values == '(1 . (2 . ()))
	ret := lang.Apply(lang.Plus, values)

	fmt.Println(ret(1))
}
