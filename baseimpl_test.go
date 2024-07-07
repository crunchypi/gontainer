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
