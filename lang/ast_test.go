package lang

import "testing"

func TestAST(t *testing.T) {
	ast := NewAST()

	if len(ast) != 2 {
		t.Errorf("S-Expressions length is wrong: %d", len(ast))
		return
	}

	if ast[0] != nil || ast[1] != nil {
		t.Errorf("Values of empty S-Expression is invalid: %s", ast)
		return
	}
}
