package lang

import "fmt"

var If λ = func(test λS) λS {
	var ret λ = func(onTrue λS) λS {
		var fret λ = func(onFalse λS) λS {
			switch test.(type) {
			case λ:
				tλ := test.(λ)
				tλtrue := tλ(onTrue.(λ)).(λ)
				return tλtrue(onFalse)
			default:
				panic(fmt.Sprintf("If expect a λ test function: %+v", test))
			}
		}

		return fret
	}

	return ret
}
