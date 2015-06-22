package lang

var True λ = func(onTrue λS) λS {
	var ret λ = func(onFalse λS) λS {
		return onTrue.(λ)(func(x λS) λS {
			return x
		})
	}

	return ret
}

var False λ = func(onTrue λS) λS {
	var ret λ = func(onFalse λS) λS {
		return onFalse.(λ)(func(x λS) λS {
			return x
		})
	}

	return ret
}
