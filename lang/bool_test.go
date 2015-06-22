package lang

import "testing"

var (
	// Zero applies the f function zero times
	Zero λ = func(f λS) λS {
		return func(x λS) λS {
			return x
		}
	}

	// Add1 creates the successive numeral
	//	Add1 λ = func(n λS) λS {
	//		return func(f λS) λS {
	//			return func(x λS) λS {
	//				return f.(λ)((n.(λ)(f)(x)))
	//			}
	//		}
	//	}

	// One is a lambda that when invoked applies the lambda f only
	// one time.
	One λ = func(f λS) λS {
		return func(x λS) λS {
			return f.(λ)(x)
		}
	}
)

func TestTrue(t *testing.T) {
	//	If(True).(λ)(func(x λS) λS {
	//		return Nil
	//	}).(λ)(func(x λS) λS {
	//		t.Error("should be true")
	//		return Nil
	//	})

	//	Add1(One).(λ)(func(x λS) λS {
	//		fmt.Println("ONE", x)
	//		return x
	//	}).(λ)(Nil)
}
