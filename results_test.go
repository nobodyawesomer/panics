package results_test

import (
	"errors"
	"strconv"
	"testing"

	. "github.com/nobodyawesomer/results"
)

func TestTry(t *testing.T) {
	errant := Try(strconv.Atoi("1x0"))
	// t.Logf("Errant: %v", errant)

	valueIfSyntaxError := 100
	errant.Catch(strconv.ErrSyntax, func(r *int, err error) {
		t.Logf("Int: %v, Err: %v", r, err)
		if !errors.Is(err, strconv.ErrSyntax) {
			t.Error("Err was supposed to be", strconv.ErrSyntax, "but err was", err)
		}
		*r = valueIfSyntaxError
	})
	t.Logf("Errant: %v", errant)
	result := errant.Unwrap()
	if result != valueIfSyntaxError {
		t.Fail()
	}
}

func TestUnwrap_Succeed(t *testing.T) {
	defer failPanicked(t)
	result := Unwrap(strconv.Atoi("25"))
	if result != 25 {
		t.Error("Result was not correct")
	}
}

func TestUnwrap_Panic(t *testing.T) {
	defer failUnpanicked(t, "")
	Unwrap(strconv.Atoi("hello"))
}

// usage: defer failPanicked(t)
func failPanicked(t *testing.T) {
	if r := recover(); r != nil {
		t.Errorf("Was not supposed to panic but panicked with '%v'", r)
	}
}

// usage: defer failUnpanicked(t, "expected").
// use "" if you have no expectations
func failUnpanicked(t *testing.T, expectedMessage string) { // TODO: convert to var-args array of expectedMessages, being the panic chain
	r := recover()
	if r == nil {
		t.Error("Was supposed to panic but didn't.")
	} else if r != expectedMessage && expectedMessage != "" {
		t.Errorf("Was supposed to panic with '%v' but instead panicked with '%v'",
			expectedMessage,
			r,
		)
	} else {
		t.Logf("Successfully panicked with expected '%v'", r)
	}
}
