package gontainer

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func assertEq[T, U any](subject string, a T, b U, f func(string)) {
	if f == nil {
		return
	}

	ab, _ := json.Marshal(a)
	bb, _ := json.Marshal(b)

	as := string(ab)
	bs := string(bb)

	if as == bs {
		return
	}

	s := "unexpected '%v':\n\twant: '%v'\n\thave: '%v'\n"
	f(fmt.Sprintf(s, subject, as, bs))
}

// -----------------------------------------------------------------------------
// Tests for PutterImpl.
// -----------------------------------------------------------------------------

func TestPutterImplIdeal(t *testing.T) {
	p := PutterImpl[int, int]{}
	p.Impl = func(context.Context, int, int) error { return nil }

	want := *new(error)
	have := p.Put(nil, 0, 0)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

func TestPutterImplWithNil(t *testing.T) {
	p := PutterImpl[int, int]{}

	want := ErrImpl
	have := p.Put(nil, 1, 1)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for GetterImpl.
// -----------------------------------------------------------------------------

func TestGetterImplIdeal(t *testing.T) {
	g := GetterImpl[int, int]{}
	g.Impl = func(context.Context, int) (r int, err error) { r = 1; return }

	val, err := g.Get(nil, 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })

}

func TestGetterImplWithNil(t *testing.T) {
	g := GetterImpl[int, int]{}

	want := ErrImpl
	_, have := g.Get(nil, 0)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for ModifierImpl.
// -----------------------------------------------------------------------------

func TestModifierImplIdeal(t *testing.T) {
	m := ModifierImpl[int, int]{}
	m.Impl = func(context.Context, int, func(int) int) error { return nil }

	want := *new(error)
	have := m.Mod(nil, 0, nil)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

func TestModifierImplWithNil(t *testing.T) {
	m := ModifierImpl[int, int]{}

	want := ErrMod
	have := m.Mod(nil, 0, nil)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for DeleterImpl.
// -----------------------------------------------------------------------------

func TestDeleterImplIdeal(t *testing.T) {
	d := DeleterImpl[int, int]{}
	d.Impl = func(context.Context, int) (r int, err error) { r = 1; return }

	val, err := d.Del(nil, 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })
}

func TestDeleterImplWithNil(t *testing.T) {
	d := DeleterImpl[int, int]{}

	want := ErrImpl
	_, have := d.Del(nil, 0)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

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
