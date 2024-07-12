package gontainer

import "testing"

func TestNewPut(t *testing.T) {
	cnt := New[int, int]()
	err := cnt.Put(nil, 1, 1)

	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
}
