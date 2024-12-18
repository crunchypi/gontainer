package gontainer

import (
	"context"
	"testing"
)

// -----------------------------------------------------------------------------
// Tests for ContainerImpl.
// -----------------------------------------------------------------------------

func TestContainerImplLenIdeal(t *testing.T) {
	c := ContainerImpl[int, int]{}
	c.ImplLen = func(context.Context) (n int, err error) { n = 1; return }

	n, err := c.Len(nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 1, n, func(s string) { t.Fatal(s) })
}

func TestContainerImplLenWithErr(t *testing.T) {
	c := ContainerImpl[int, int]{}

	want := ErrImpl
	_, have := c.Len(nil)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

func TestContainerImplCapIdeal(t *testing.T) {
	c := ContainerImpl[int, int]{}
	c.ImplCap = func(context.Context) (n int, err error) { n = 1; return }

	n, err := c.Cap(nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("cap", 1, n, func(s string) { t.Fatal(s) })
}

func TestContainerImplCapWithErr(t *testing.T) {
	c := ContainerImpl[int, int]{}

	want := ErrImpl
	_, have := c.Cap(nil)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}
