package lang

import (
	"fmt"
	"testing"
)

func TestSExprs(t *testing.T) {
	nilValue := Nil(1)

	if nilValue != nil {
		t.Errorf("Nil isn't nil: %+v", nilValue)
		return
	}

	atom1 := NewAtom(1)

	if atom1 == nil {
		t.Error("Failed to init atom")
		return
	}

	atomValue := atom1(1)

	if atomValue != 1 {
		t.Errorf("Invalid atom value: %d", atomValue)
		return
	}

	atomInvalid := atom1(2)

	if atomInvalid != nil {
		t.Errorf("Is an atom, not s-exprs: %+v", atomInvalid)
		return
	}

	atomRep := atom1.ConsString()

	if atomRep != "(1 . ())" {
		t.Errorf("Invalid atom representation: %s", atomRep)
		return
	}

	sexprs := Cons(atom1, Nil)

	sexprsRep := sexprs.ConsString()

	if sexprsRep != "(1 . ())" {
		t.Errorf("Invalid sexpression representation: %s", sexprsRep)
		return
	}

	sexprs = Cons(NewAtom(2), NewAtom(sexprs))
	sexprsRep = sexprs.ConsString()
	if sexprsRep != "(2 . (1 . ()))" {
		t.Errorf("Invalid sexpression representation: %s", sexprsRep)
		return
	}

	sexprs = Cons(NewAtom("teste"), NewAtom(sexprs))
	sexprsRep = sexprs.ConsString()
	if sexprsRep != `("teste" . (2 . (1 . ())))` {
		t.Errorf("Invalid sexpression representation: %s", sexprsRep)
		return
	}

	sexprs = Cons(NewAtom(Cons(NewAtom("ENZO"), Nil)), Nil)
	sexprsRep = sexprs.ConsString()
	if sexprsRep != `(("ENZO" . ()) . ())` {
		t.Errorf("Invalid sexpression representation: %s", sexprsRep)
		return
	}
}

func TestAST(t *testing.T) {
	value2 := Cons(func(uint) interface{} { return 2 }, func(uint) interface{} { return nil })
	value1 := Cons(func(uint) interface{} { return 1 }, func(uint) interface{} { return value2 })

	fmt.Println("ToString: ", value2.ConsString())

	Apply(Print, value2)

	ret := Apply(Plus, value1)

	fmt.Println(ret(1))
}
