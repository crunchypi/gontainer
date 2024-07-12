package gontainer

import "testing"

func TestNewPut(t *testing.T) {
	cnt := New[int, int]()
	err := cnt.Put(nil, 1, 1)

	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
}

func TestNewGet(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	val := 0

	// First call to Get should return an err.
	val, err = cnt.Get(nil, 1)
	assertEq("err", ErrGet, err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(nil, 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Get again to validate the placement.
	val, err = cnt.Get(nil, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })
}
