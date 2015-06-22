package lang

import "fmt"

var Print λ = func(v λS) λS {
	switch v.(type) {
	case λ:
		fmt.Println(v.(λ)(True))
	default:
		fmt.Println(v)
	}

	return True
}
