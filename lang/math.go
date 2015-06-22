package lang

// Plus is a recursive +
func Plus(exprs λS) λS {
	values := exprs.(SExprs)
	var consS = Env["cons"]

	if consS == nil {
		panic("missing core library")
		return Nil
	}

	cons := consS(1).(λ)

	if cons == nil {
		panic("Invalid cons procedure type")
		return Nil
	}

	cdr, ok := values(False).(SExprs)

	if !ok {
		return NewAtom(values(True))
	}

	car, ok := values(True).(int)

	if !ok {
		return Nil
	}

	cdrcar, ok := cdr(True).(int)

	if !ok {
		return Nil
	}

	acc := cdrcar + car

	cdrcdr, ok := cdr(False).(SExprs)

	if !ok {
		return cons(NewAtom(acc)).(λ)(Nil)
	}

	return Plus(cons(NewAtom(acc)).(λ)(cdrcdr))
}
