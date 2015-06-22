// Package lang has the core concepts of Scherzo
//
// Axiom's:
//
// I- S-Expression is a lambda that can return an atom or a S-Expression;
// II- Atom is a lambda that returns a symbol or literal;
// III- Symbol is anything named in the program;

package lang

import "fmt"

type λS interface{}

type λ func(λS) λS

var (
	Nil λ = func(i λS) λS {
		return nil
	}
)

// SExprs in Scherzo is defined as:
//     - An λ that closures x and y and have the form:
//         (λ (pick)
//            (cond ((= pick True) x)
//                  ((= pick False) y)))
type SExprs λ

func (s SExprs) ConsString() string {
	var (
		value interface{}
	)

	toString := func(pick λS) string {
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

	repCar := toString(True)
	repCdr := toString(False)

	if repCdr == "" {
		return repCar
	}

	return "(" + repCar + " . " + repCdr + ")"
}

// Atom is an S-Expression that returns an literal when evaluated
// Atom haven't an explicit type

// NewAtom returns a new Atom
func NewAtom(value interface{}) SExprs {
	var ret SExprs = func(i λS) λS {
		// If(True)
		fmt.Println(i)
		var vλ λ = func(x λS) λS {
			return value.(λS)
		}

		ifV := λ(If(i.(λ)).(λ))(vλ)

		// If(False)
		var ifFalse λ = func(x λS) λS {
			return λS(Nil)
		}

		return ifV.(λ)(ifFalse)
	}

	return ret
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
	var ret λ = func(b λS) λS {
		var ret2 SExprs = func(pick λS) λS {
			var onTrue λ = func(x λS) λS {
				return a
			}

			var onFalse = func(x λS) λS {
				return b
			}

			return If(pick).(λ)(onTrue).(λ)(onFalse)
		}

		return ret2
	}

	return ret
}

func Define(name string, value SExprs) SExprs {
	Env[name] = value
	return value
}

func Apply(operator λ, operands λS) λS {
	return operator(operands)
}
