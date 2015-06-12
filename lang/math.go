package lang

// Plus is a recursive +
func Plus(exprs ...SExprs) SExprs {
	if len(exprs) > 1 {
		return Nil
	}

	values := exprs[0]

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

	cdr, ok := values(2).(SExprs)

	if !ok {
		return NewAtom(values(1))
	}

	car, ok := values(1).(int)

	if !ok {
		return Nil
	}

	cdrcar, ok := cdr(1).(int)

	if !ok {
		return Nil
	}

	acc := cdrcar + car

	cdrcdr, ok := cdr(2).(SExprs)

	if !ok {
		return cons(NewAtom(acc), Nil)
	}

	return Plus(cons(NewAtom(acc), cdrcdr))
}
