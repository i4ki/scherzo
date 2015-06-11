package lang

type λ func(SExprs) SExprs

// SExprs in Scherzo is defined as:
//     - An λ that closures x and y and have the form:
//         (λ (pick)
//            (cond ((= pick 1) x)
//                  ((= pick 2) y)))
type SExprs func(uint) interface{}

func NewAtom(value interface{}) SExprs {
	return func(uint) interface{} { return value }
}

var Nil = NewAtom(nil)

// Temporary Cons implementation. This should be defined in Scherzo user library
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

// Plus is a recursive +
func Plus(values SExprs) SExprs {
	cdr, ok := values(2).(SExprs)

	if !ok {
		return func(uint) interface{} {
			return values(1)
		}
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
		return Cons(NewAtom(acc), Nil)
	}

	return Plus(Cons(NewAtom(acc), cdrcdr))
}

func Apply(operator λ, operands SExprs) SExprs {
	return operator(operands)
}
