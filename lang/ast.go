package lang

import "errors"

type Element interface {
	// S-Exprs OR Atom
}

type Atom interface{}

type Cons struct {
	Head Element
	Tail Element
}

type SExprs [2]Element

var Nil = SExprs{}

func NewSExprs() SExprs {
	ast := Nil

	return ast
}

func (l *SExprs) Pick(elem uint) (Element, error) {
	if elem > 1 {
		return nil, errors.New("S-Exprs is a pair. Pick can receive only 0 or 1")
	}
}
