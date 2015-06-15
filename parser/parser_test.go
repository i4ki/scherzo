package parser

import (
	"testing"

	"github.com/juju/testing/checkers"
	"github.com/tiago4orion/scherzo/lang"
)

func TestSimpleExpression(t *testing.T) {
	src := `(+ (1 (2 ())))`

	exprs, err := FromString(src)

	if err != nil {
		t.Error(err)
		return
	}

	two := lang.NewAtom(2)
	tail := lang.Cons(two, lang.Nil)
	head := lang.Cons(lang.NewAtom(1), lang.NewAtom(tail))
	proc := lang.Cons(lang.NewAtom("+"), lang.NewAtom(head))

	ok, err := checkers.DeepEqual(proc.ConsString(), exprs.ConsString())

	if !ok || err != nil {
		t.Errorf("Parsed differs: %s != %s", proc.ConsString(), exprs.ConsString())
		t.Error(err.Error())
	}
}
