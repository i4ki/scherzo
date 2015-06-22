package lang

func If(test λ) λ {
	return func(onTrue λ) λ {
		return func(onFalse λ) λ {
			return test(onTrue)(onFalse)
		}
	}
}
