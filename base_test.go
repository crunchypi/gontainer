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

func TestNewMod(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	val := 0

	// First mod should upsert, in this case the zero-value + 1.
	err = cnt.Mod(nil, 1, func(v int) int { return v + 1 })
	assertEq("err", ErrMod, err, func(s string) { t.Fatal(s) })

	// Validate that the value is there.
	val, err = cnt.Get(nil, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })

	// Mod again by incrementing the first value.
	err = cnt.Mod(nil, 1, func(v int) int { return v + 1 })
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Get again, validate the increment.
	val, err = cnt.Get(nil, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 2, val, func(s string) { t.Fatal(s) })

	// Mod with nil func.
	err = cnt.Mod(nil, 1, nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Get to validate that nothing has changed with the key.
	val, err = cnt.Get(nil, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 2, val, func(s string) { t.Fatal(s) })
}

func TestNewDel(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	val := 0

	// First Del should return an err since nothing was deleted.
	val, err = cnt.Del(nil, 1)
	assertEq("err", ErrDel, err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(nil, 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Now Del should return the added value.
	val, err = cnt.Del(nil, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })

	// Validate that the value was deleted.
	val, err = cnt.Get(nil, 1)
	assertEq("err", ErrGet, err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestNewLen(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	l := 0

	// First Len call should return 0.
	l, err = cnt.Len(nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 0, l, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(nil, 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Now Len should return 1.
	l, err = cnt.Len(nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 1, l, func(s string) { t.Fatal(s) })
}

func TestNewCap(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	c := 0

	// First Cap call should return 0.
	c, err = cnt.Cap(nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 0, c, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(nil, 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Now Cap should return double len.
	c, err = cnt.Cap(nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 2, c, func(s string) { t.Fatal(s) })
}
