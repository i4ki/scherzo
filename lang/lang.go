// Package lang has the core concepts of Scherzo
//
// Axiom's:
//
// I- S-Expression is a lambda that can return an atom or a S-Expression;
// II- Atom is a lambda that returns a symbol or literal;
// III- Symbol is anything named in the program;

package lang

type λ func(...SExprs) SExprs

// SExprs in Scherzo is defined as:
//     - An λ that closures x and y and have the form:
//         (λ (pick)
//            (cond ((= pick 1) x)
//                  ((= pick 2) y)))
type SExprs func(uint) interface{}

// Atom is an S-Expression that returns an literal when evaluated
// Doesn't have an explicit type

// NewAtom returns a new Atom
func NewAtom(value interface{}) SExprs {
	return func(uint) interface{} { return value }
}

var Nil = NewAtom(nil)
var Env = map[string]SExprs{}

func init() {
	var cons = λ(Cons)

	Env["cons"] = NewAtom(cons)
	Env["print"] = NewAtom(Print)
}

// Cons is the List constructor.
// Temporary Cons implementation. This should be defined in scherzo lang.
func Cons(ab ...SExprs) SExprs {
	if len(ab) != 2 {
		panic("invalid cons implementation")
	}

	a := ab[0]
	b := ab[1]

	return func(pick uint) interface{} {
		if pick == 1 {
			return a(1)
		} else if pick == 2 {
			return b(1)
		}

		return nil
	}
}

func Define(name string, value SExprs) SExprs {
	Env[name] = value
	return value
}

func Apply(operator λ, operands SExprs) SExprs {
	return operator(operands)
}
