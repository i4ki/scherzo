package lang

import (
	"fmt"
	"testing"
)

func TestSExprs(t *testing.T) {
	nilValue := Nil(True)

	if nilValue != nil {
		t.Errorf("Nil isn't nil: %+v", nilValue)
		return
	}

	atom1 := NewAtom(1)

	if atom1 == nil {
		t.Error("Failed to init atom")
		return
	}

	atomValue := atom1(True)

	if atomValue != 1 {
		t.Errorf("Invalid atom value: %d", atomValue)
		return
	}

	atomInvalid := atom1(False)

	if atomInvalid != 1 {
		switch atomInvalid.(type) {
		case λS:
			val := atomInvalid.(λ)(Nil)
			if val != nil {
				t.Errorf("Is an atom, not s-exprs: %+v", val)
			}
		default:
			t.Errorf("Is an atom, not s-exprs: %+v", atomInvalid)
		}
	}

	//atomRep := atom1.ConsString()

	//if atomRep != "(1 . ())" {
	//	t.Errorf("Invalid atom representation: %s", atomRep)
	//	return
	//}

	//sexprs := Cons(atom1, Nil)

	//sexprsRep := sexprs.ConsString()

	//if sexprsRep != "(1 . ())" {
	//	t.Errorf("Invalid sexpression representation: %s", sexprsRep)
	//	return
	//}

	//sexprs = Cons(NewAtom(2), NewAtom(sexprs))
	//sexprsRep = sexprs.ConsString()
	//if sexprsRep != "(2 . (1 . ()))" {
	//	t.Errorf("Invalid sexpression representation: %s", sexprsRep)
	//	return
	//}

	//sexprs = Cons(NewAtom("teste"), NewAtom(sexprs))
	//sexprsRep = sexprs.ConsString()
	//if sexprsRep != `("teste" . (2 . (1 . ())))` {
	//	t.Errorf("Invalid sexpression representation: %s", sexprsRep)
	//	return
	//}

	//sexprs = Cons(NewAtom(Cons(NewAtom("ENZO"), Nil)), Nil)
	//sexprsRep = sexprs.ConsString()
	//if sexprsRep != `(("ENZO" . ()) . ())` {
	//	t.Errorf("Invalid sexpression representation: %s", sexprsRep)
	//	return
	//}
}

func TestAST(t *testing.T) {
	var vv λ = Cons(NewAtom(2)).(λ)

	value2 := vv(NewAtom(nil)).(SExprs)
	var vv2 λ = Cons(NewAtom(1)).(λ)

	value1 := vv2(NewAtom(value2))

	fmt.Println("ToString: ", value2.ConsString())

	Apply(Print, value2)

	ret := Apply(Plus, value1)

	fmt.Println(ret.(SExprs)(True))
}
