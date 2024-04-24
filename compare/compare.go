package compare

func Compare[N Ordered](a N, b N) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return +1
	default:
		return 0
	}
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T Ordered](x T) bool {
	return x != x
}

// Identity is the identity function for something that is comparable
func Identity[N comparable](n N) N {
	return n
}

// Equal is the equality function for something that is comparable
func Equal[N comparable](a, b N) bool {
	return a == b
}

// Less returns true if a < b
func Less[E Ordered](a, b E) bool {
	return a < b
}

// LessOrEqual returns true if a <= b
func LessOrEqual[E Ordered](a, b E) bool {
	return a <= b
}

// Negate will return a function negating the result from other function. The use case for this is reversing a comparison
func Negate[A any](f func(a, b A) bool) func(A, A) bool {
	return func(a, b A) bool {
		return !f(a, b)
	}
}

// EqualOf returns a function that compares the input to the original input
func EqualOf[N comparable](needle N) func(b N) bool {
	return func(b N) bool {
		return needle == b
	}
}

// IsZero returns function that looks at the input that returns true if passed in argument is zero
func IsZero[N comparable]() func(b N) bool {
	var n N
	return EqualOf(n)
}

// IsNotZero returns function that looks at the input that returns true if passed in argument is *Not* zero
func IsNotZero[N comparable]() func(b N) bool {
	return NegateOf(IsZero[N]())
}

// NegateOf returns a function the negates the result of the original function passed in
func NegateOf[A any](f func(A) bool) func(A) bool {
	return func(a A) bool {
		return !f(a)
	}
}

// Ternary is equivalent to "expression ? a : b" ternary notation and returns ifTrue if true and ifFalse if false
func Ternary[A any](boolean bool, ifTrue A, ifFalse A) A {
	if boolean {
		return ifTrue
	}
	return ifFalse
}

// Coalesce returns the first of its arguments that is not equal to the zero value.
// If no argument is non-zero, it returns the zero value.
func Coalesce[T comparable](vals ...T) (t T) {
	zero := IsZero[T]()
	for _, val := range vals {
		if !zero(val) {
			return val
		}
	}
	return t
}
