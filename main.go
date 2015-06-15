package main

import (
	"fmt"
	"os"

	"github.com/tiago4orion/scherzo/lang"
	"github.com/tiago4orion/scherzo/parser"
)

const version = "0.0.1"

func main() {
	fmt.Println("Scherzo compiler ", version)

	file, err := os.Open("./hello.scm")

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	_, err = parser.FromReader(file)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	Example()
}

// Example represents the AST of the code below:
//
// (print (+ 1 2))
func Example() {
	value2 := lang.Cons(lang.NewAtom(2), lang.Nil)
	values := lang.Cons(lang.NewAtom(1), lang.NewAtom(value2))

	// values == '(1 . (2 . ()))
	ret := lang.Apply(lang.Plus, values)
	lang.Apply(lang.Print, ret)
}
