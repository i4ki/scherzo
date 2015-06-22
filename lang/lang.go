// Package lang has the core concepts of Scherzo
//
// Axiom's:
//
// I- S-Expression is a lambda that can return an atom or a S-Expression;
// II- Atom is a lambda that returns a symbol or literal;
// III- Symbol is anything named in the program;

package lang

type λS interface{}

type λ func(λS) λS

var (
	Nil λ = func(i λS) λS {
		return i
	}
)

// SExprs in Scherzo is defined as:
//     - An λ that closures x and y and have the form:
//         (λ (pick)
//            (cond ((= pick True) x)
//                  ((= pick False) y)))
type SExprs λS

// Atom is an S-Expression that returns an literal when evaluated
// Atom haven't an explicit type

// NewAtom returns a new Atom
func NewAtom(value interface{}) SExprs {
	return func(i λ) interface{} {
		return If(i)(func(x λ) λS {
			return value.(λS)
		}).(λ)(func(x λ) interface{} {
			return λS(Nil)
		})
	}
}

var Env = map[string]SExprs{}

func init() {
	var cons = λ(Cons)

	Env["cons"] = NewAtom(cons)
	Env["print"] = NewAtom(Print)
}

// Cons is the List constructor.
// Temporary Cons implementation. This should be defined in scherzo lang.
func Cons(a λS) λS {
	return func(b λS) λS {
		return func(pick λ) λS {
			return If(pick)(func(x λS) λS {
				return a
			}).(λ)(func(x λS) λS {
				return b
			})
		}
	}
}

func Define(name string, value SExprs) SExprs {
	Env[name] = value
	return value
}

func Apply(operator λ, operands SExprs) SExprs {
	return operator(operands)
}
