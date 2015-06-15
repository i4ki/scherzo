// Package lang has the core concepts of Scherzo
//
// Axiom's:
//
// I- S-Expression is a lambda that can return an atom or a S-Expression;
// II- Atom is a lambda that returns a symbol or literal;
// III- Symbol is anything named in the program;

package lang

import "fmt"

type λ func(...SExprs) SExprs

// SExprs in Scherzo is defined as:
//     - An λ that closures x and y and have the form:
//         (λ (pick)
//            (cond ((= pick 1) x)
//                  ((= pick 2) y)))
type SExprs func(uint) interface{}

func (s SExprs) ConsString() string {
	var (
		value interface{}
	)

	toString := func(pick uint) string {
		value = s(pick)
		switch value.(type) {
		case nil:
			return "()"
		case int:
			return fmt.Sprintf("%d", value)
		case bool:
			boolValue := value.(bool)
			if boolValue {
				return "true"
			} else {
				return "false"
			}
		case string:
			return fmt.Sprintf(`"%s"`, value.(string))
		case SExprs:
			vsexpr := value.(SExprs)
			return vsexpr.ConsString()
		default:
			panic(fmt.Sprintf("invalid sexpression: %s", value))
		}

		return ""
	}

	repCar := toString(1)
	repCdr := toString(2)

	if repCdr == "" {
		return repCar
	}

	return "(" + repCar + " . " + repCdr + ")"
}

// Atom is an S-Expression that returns an literal when evaluated
// Atom haven't an explicit type

// NewAtom returns a new Atom
func NewAtom(value interface{}) SExprs {
	return func(i uint) interface{} {
		if i == 2 {
			return nil
		}

		return value
	}
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
