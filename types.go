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

	// UnwrapOrElse attempts to resolve the Result.
	// If there is an error, it runs the onError callback to
	// generate a new value.
	// Otherwise, it returns its underlying value.
	UnwrapOrElse(onError func() R) R

	// Catch attempts to catch a particular error type.
	// If the error is a match for a given errorType, then it runs
	// the onCatch callback and sets the error type to nil.
	// Otherwise, the error, if it exists, is passed through along
	// with the value.
	// The Result is passed back for chaining convenience.
	Catch(errorType error, onCatch func(*R, error)) Result[R]

	// CatchTop attempts to unwraps one layer off the error chain and
	// "catch" (process) it according to the provided
	// onCatch callback.
	// If there is no error, the callback is not run and the
	// Result is passed along unchanged.
	// Otherwise, the resulting error, if it exists,
	// is then passed through along with the value.
	// The Result is passed back for chaining convenience.
	//
	// Note that the full error chain is provided to onCatch, not
	// just the top error.
	CatchTop(onCatch func(*R, error)) Result[R]

	// CatchAll catches all remaining errors and runs the provided onCatch
	// callback. It then returns the underlying value.
	CatchAll(onCatch func(*R, error)) R
}
