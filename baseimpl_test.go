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

// -----------------------------------------------------------------------------
// Tests for SearcherImpl.
// -----------------------------------------------------------------------------

func TestSearcherImplIdeal(t *testing.T) {
	s := SearcherImpl[int, int]{}
	s.Impl = func(_ context.Context, _ int) (r int, err error) { return }

	val, err := s.Search(nil, 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestSearcherImplWithNil(t *testing.T) {
	s := SearcherImpl[int, int]{}

	_, err := s.Search(nil, 0)
	assertEq("err", ErrImpl, err, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for SearchUpdaterImpl.
// -----------------------------------------------------------------------------

func TestSearchUpdaterImplIdeal(t *testing.T) {
	s := SearchUpdaterImpl[int, int, int]{}
	s.Impl = func(_ context.Context, _, _ int) (r int, err error) { return }

	val, err := s.SearchUpdate(nil, 0, 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestSearchUpdaterImplWithNil(t *testing.T) {
	s := SearchUpdaterImpl[int, int, int]{}

	_, err := s.SearchUpdate(nil, 0, 0)
	assertEq("err", ErrImpl, err, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for SearchDeleterImpl.
// -----------------------------------------------------------------------------

func TestSearchDeleterImplIdeal(t *testing.T) {
	s := SearchDeleterImpl[int, int]{}
	s.Impl = func(_ context.Context, _ int) (q int, err error) { return }

	val, err := s.SearchDelete(nil, 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestSearchDeleterImplWithNil(t *testing.T) {
	s := SearchDeleterImpl[int, int]{}

	_, err := s.SearchDelete(nil, 0)
	assertEq("err", ErrImpl, err, func(s string) { t.Fatal(s) })
}
