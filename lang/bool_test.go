package lang

import (
	"fmt"
	"testing"
)

var (
	Void λ = func(x λ) λ {
		return x
	}

	True λ = func(onTrue λ) λ {
		return func(onFalse λ) λ {
			return onTrue(Void)
		}
	}

	False λ = func(onTrue λ) λ {
		return func(onFalse λ) λ {
			return onTrue(Void)
		}
	}

	// Zero applies the f function zero times
	Zero λ = func(f λ) λ {
		return func(x λ) λ {
			return x
		}
	}

	// Add1 creates the successive numeral
	Add1 λ = func(n λ) λ {
		return func(f λ) λ {
			return func(x λ) λ {
				return f((n(f)(x)))
			}
		}
	}

	// One is a lambda that when invoked applies the lambda f only
	// one time.
	One λ = func(f λ) λ {
		return func(x λ) λ {
			return f(x)
		}
	}
)

func TestTrue(t *testing.T) {
	If(True)(func(x λ) λ {
		return Void
	})(func(x λ) λ {
		t.Error("should be true")
		return Void
	})

	Add1(One)(func(x λ) λ {
		fmt.Println("ONE", x)
		return x
	})(Void)

}
