package lang

type Element interface {
	// S-Exprs OR Atom
}

type Atom interface{}

type λ func(SExprs) SExprs

// SExprs in Scherzo is defined as:
//     - An λ that closures x and y and have the form:
//         (λ (pick)
//            (cond ((= pick 1) x)
//                  ((= pick 2) y)))
type SExprs func(uint) interface{}

// Should be defined in Scherzo user library
func Cons(a SExprs, b SExprs) SExprs {
	return func(pick uint) interface{} {
		if pick == 1 {
			return a(1)
		} else if pick == 2 {
			return b(1)
		}

		return nil
	}
}

func Plus(values SExprs) SExprs {
	cdr, ok := values(2).(SExprs)

	if !ok {
		return func(uint) interface{} {
			return values(1)
		}
	}

	car, ok := values(1).(int)

	if !ok {
		return func(uint) interface{} {
			return nil
		}
	}

	cdrcar, ok := cdr(1).(int)

	if !ok {
		return func(uint) interface{} {
			return nil
		}
	}

	acc := cdrcar + car

	cdrcdr, ok := cdr(2).(SExprs)

	if !ok {
		return Cons(func(uint) interface{} {
			return acc
		}, func(uint) interface{} {
			return nil
		})
	}

	return Plus(Cons(func(uint) interface{} {
		return cdrcar + car
	}, cdrcdr))
}

func Apply(operator λ, operands SExprs) SExprs {
	return operator(operands)
}
