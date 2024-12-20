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
	have := p.Put(context.Background(), 0, 0)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

func TestPutterImplWithNil(t *testing.T) {
	p := PutterImpl[int, int]{}

	want := ErrImpl
	have := p.Put(context.Background(), 1, 1)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for GetterImpl.
// -----------------------------------------------------------------------------

func TestGetterImplIdeal(t *testing.T) {
	g := GetterImpl[int, int]{}
	g.Impl = func(context.Context, int) (r int, err error) { r = 1; return }

	val, err := g.Get(context.Background(), 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })

}

func TestGetterImplWithNil(t *testing.T) {
	g := GetterImpl[int, int]{}

	want := ErrImpl
	_, have := g.Get(context.Background(), 0)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for ModifierImpl.
// -----------------------------------------------------------------------------

func TestModifierImplIdeal(t *testing.T) {
	m := ModifierImpl[int, int]{}
	m.Impl = func(context.Context, int, func(int) int) error { return nil }

	want := *new(error)
	have := m.Mod(context.Background(), 0, nil)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

func TestModifierImplWithNil(t *testing.T) {
	m := ModifierImpl[int, int]{}

	want := ErrMod
	have := m.Mod(context.Background(), 0, nil)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for DeleterImpl.
// -----------------------------------------------------------------------------

func TestDeleterImplIdeal(t *testing.T) {
	d := DeleterImpl[int, int]{}
	d.Impl = func(context.Context, int) (r int, err error) { r = 1; return }

	val, err := d.Del(context.Background(), 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })
}

func TestDeleterImplWithNil(t *testing.T) {
	d := DeleterImpl[int, int]{}

	want := ErrImpl
	_, have := d.Del(context.Background(), 0)
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for SearcherImpl.
// -----------------------------------------------------------------------------

func TestSearcherImplIdeal(t *testing.T) {
	s := SearcherImpl[int, int]{}
	s.Impl = func(_ context.Context, _ int) (r int, err error) { return }

	val, err := s.Search(context.Background(), 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestSearcherImplWithNil(t *testing.T) {
	s := SearcherImpl[int, int]{}

	_, err := s.Search(context.Background(), 0)
	assertEq("err", ErrImpl, err, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for SearchUpdaterImpl.
// -----------------------------------------------------------------------------

func TestSearchUpdaterImplIdeal(t *testing.T) {
	s := SearchUpdaterImpl[int, int, int]{}
	s.Impl = func(_ context.Context, _, _ int) (r int, err error) { return }

	val, err := s.SearchUpdate(context.Background(), 0, 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestSearchUpdaterImplWithNil(t *testing.T) {
	s := SearchUpdaterImpl[int, int, int]{}

	_, err := s.SearchUpdate(context.Background(), 0, 0)
	assertEq("err", ErrImpl, err, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for SearchDeleterImpl.
// -----------------------------------------------------------------------------

func TestSearchDeleterImplIdeal(t *testing.T) {
	s := SearchDeleterImpl[int, int]{}
	s.Impl = func(_ context.Context, _ int) (q int, err error) { return }

	val, err := s.SearchDelete(context.Background(), 0)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestSearchDeleterImplWithNil(t *testing.T) {
	s := SearchDeleterImpl[int, int]{}

	_, err := s.SearchDelete(context.Background(), 0)
	assertEq("err", ErrImpl, err, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Tests for ContainerImpl.
// -----------------------------------------------------------------------------

func TestContainerImplLenIdeal(t *testing.T) {
	c := ContainerImpl[int, int]{}
	c.ImplLen = func(context.Context) (n int, err error) { n = 1; return }

	n, err := c.Len(context.Background())
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 1, n, func(s string) { t.Fatal(s) })
}

func TestContainerImplLenWithErr(t *testing.T) {
	c := ContainerImpl[int, int]{}

	want := ErrImpl
	_, have := c.Len(context.Background())
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

func TestContainerImplCapIdeal(t *testing.T) {
	c := ContainerImpl[int, int]{}
	c.ImplCap = func(context.Context) (n int, err error) { n = 1; return }

	n, err := c.Cap(context.Background())
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("cap", 1, n, func(s string) { t.Fatal(s) })
}

func TestContainerImplCapWithErr(t *testing.T) {
	c := ContainerImpl[int, int]{}

	want := ErrImpl
	_, have := c.Cap(context.Background())
	assertEq("err", want, have, func(s string) { t.Fatal(s) })
}

// -----------------------------------------------------------------------------
// Map.
// -----------------------------------------------------------------------------

func TestNewPut(t *testing.T) {
	cnt := New[int, int]()
	err := cnt.Put(context.Background(), 1, 1)

	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
}

func TestNewGet(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	val := 0

	// First call to Get should return an err.
	val, err = cnt.Get(context.Background(), 1)
	assertEq("err", ErrGet, err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(context.Background(), 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Get again to validate the placement.
	val, err = cnt.Get(context.Background(), 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })
}

func TestNewMod(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	val := 0

	// First mod should upsert, in this case the zero-value + 1.
	err = cnt.Mod(context.Background(), 1, func(v int) int { return v + 1 })
	assertEq("err", ErrMod, err, func(s string) { t.Fatal(s) })

	// Validate that the value is there.
	val, err = cnt.Get(context.Background(), 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })

	// Mod again by incrementing the first value.
	err = cnt.Mod(context.Background(), 1, func(v int) int { return v + 1 })
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Get again, validate the increment.
	val, err = cnt.Get(context.Background(), 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 2, val, func(s string) { t.Fatal(s) })

	// Mod with nil func.
	err = cnt.Mod(context.Background(), 1, nil)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Get to validate that nothing has changed with the key.
	val, err = cnt.Get(context.Background(), 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 2, val, func(s string) { t.Fatal(s) })
}

func TestNewDel(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	val := 0

	// First Del should return an err since nothing was deleted.
	val, err = cnt.Del(context.Background(), 1)
	assertEq("err", ErrDel, err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(context.Background(), 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Now Del should return the added value.
	val, err = cnt.Del(context.Background(), 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("val", 1, val, func(s string) { t.Fatal(s) })

	// Validate that the value was deleted.
	val, err = cnt.Get(context.Background(), 1)
	assertEq("err", ErrGet, err, func(s string) { t.Fatal(s) })
	assertEq("val", 0, val, func(s string) { t.Fatal(s) })
}

func TestNewLen(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	l := 0

	// First Len call should return 0.
	l, err = cnt.Len(context.Background())
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 0, l, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(context.Background(), 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Now Len should return 1.
	l, err = cnt.Len(context.Background())
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 1, l, func(s string) { t.Fatal(s) })
}

func TestNewCap(t *testing.T) {
	cnt := New[int, int]()
	err := *new(error)
	c := 0

	// First Cap call should return 0.
	c, err = cnt.Cap(context.Background())
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 0, c, func(s string) { t.Fatal(s) })

	// Add a value.
	err = cnt.Put(context.Background(), 1, 1)
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })

	// Now Cap should return double len.
	c, err = cnt.Cap(context.Background())
	assertEq("err", *new(error), err, func(s string) { t.Fatal(s) })
	assertEq("len", 2, c, func(s string) { t.Fatal(s) })
}
