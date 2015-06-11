package main

import (
	"fmt"

	"github.com/tiago4orion/scherzo/lang"
)

const version = "0.0.1"

func main() {
	fmt.Println("Scherzo compiler ", version)

	value2 := lang.Cons(func(uint) interface{} {
		return 2
	}, func(uint) interface{} {
		return nil
	})

	values := lang.Cons(func(uint) interface{} {
		return 1
	}, func(uint) interface{} {
		return value2
	})

	// values == '(1 . (2 . ()))
	ret := lang.Apply(lang.Plus, values)

	fmt.Println(ret(1))
}
