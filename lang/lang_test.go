package lang

import (
	"fmt"
	"testing"
)

func TestAST(t *testing.T) {
	value2 := Cons(func(uint) interface{} { return 2 }, func(uint) interface{} { return nil })
	value1 := Cons(func(uint) interface{} { return 1 }, func(uint) interface{} { return value2 })

	Apply(Print, value2)

	fmt.Println(value1(1))
	fmt.Println(value1(2))
	ret := Apply(Plus, value1)

	fmt.Println(ret(1))
}
