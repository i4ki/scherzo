package lang

import "fmt"

var Print λ = func(exprs ...SExprs) SExprs {
	for _, s := range exprs {
		fmt.Println(s(1))
	}

	return Nil
}
