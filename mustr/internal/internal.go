// Copyright (c) 2025 Timo Savola
// SPDX-License-Identifier: BSD-3-Clause

package internal

// Result value with associated error.
type Result[T any] struct {
	Value T
	Err   error
}

// Testing conforms to testing.T.
type Testing interface {
	Fatal(args ...any)
	Helper()
}
