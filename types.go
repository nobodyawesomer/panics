package results

type Result[R any] interface {
	// Out outputs the Result unresolved.
	Out() (R, error)

	// Unwrap attempts to resolve the Result.
	// If there is an error, it panics.
	// Otherwise, it returns its underlying value.
	Unwrap() R

	// UnwrapOr resolves the Result safely.
	// If there is an error, it quietly returns the provided value.
	// Otherwise, it returns its underlying value.
	UnwrapOr(value R) R

	// UnwrapOrElse resolves the Result safely.
	// If there is an error, it runs the onError callback to
	// generate a new value.
	// Otherwise, it returns its underlying value.
	// onError must not be nil. If you would like to utilize the error,
	// see CatchAll.
	UnwrapOrElse(onError func() R) R

	// UnwrapOrDefault resolves the Result safely.
	// If there is an error, it generates a new value using new(R),
	// yielding the zero value for that type.
	// Otherwise, it returns its underlying value.
	UnwrapOrDefault() R

	// Catch attempts to catch a particular error type.
	// If the error is a match for a given errorType, then it runs
	// the onCatch callback and sets the error type to nil.
	// Otherwise, the error, if it exists, is passed through unchanged
	// along with the value.
	//
	// If onCatch is nil, it will catch the error type and do nothing.
	// This is equivalent to passing it an empty function.
	//
	// The Result is passed back for chaining convenience.
	Catch(errorType error, onCatch func(*R, error)) Result[R]

	// CatchAnd attempts to catch a particular error type.
	// If the error is a match for a given errorType, then it sets
	// the value to value and the error to nil.
	// Otherwise, the error, if it exists, is passed through unchanged
	// along with the value.
	//
	// The Result is passed back for chaining convenience.
	CatchAnd(errorType error, value R) Result[R]

	// CatchAll catches all remaining errors and runs the provided onCatch
	// callback. It then returns the underlying value.
	// onCatch must not be nil.
	CatchAll(onCatch func(*R, error)) R
}
