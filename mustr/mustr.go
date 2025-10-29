// Copyright (c) 2025 Timo Savola
// SPDX-License-Identifier: BSD-3-Clause

// Package mustr is guaranteed to only export symbols starting with word Must
// or R.  It is intended to be used in tests via dot-import:
//
//	import (
//	    "testing"
//
//	    . "import.name/testing/mustr"
//	)
//
//	func example() (int, error)
//
//	func Test(t *testing.T) {
//	    _ = Must(t, R(example()))
//	}
//
// Test fails immediately if example returns an error.
package mustr

import "import.name/testing/mustr/internal"

// R can be used to wrap the result of a function call which returns two
// values, the second of which is an error.
func R[Value any](x Value, err error) internal.Result[Value] {
	return internal.Result[Value]{Value: x, Err: err}
}

// Must calls t.Fatal() with the error (the second value of the result) if it
// is non-nil, otherwise it returns the first value of the result.
func Must[Value any](t internal.Testing, r internal.Result[Value]) Value {
	t.Helper()
	if r.Err != nil {
		t.Fatal(r.Err)
	}
	return r.Value
}
