package gontainer

import (
	"context"
)

// Container groups Putter, Getter, Modifier and Deleter. Additionally, it
// also defines the expectation of Len and Cap.
type Container[K comparable, V any] interface {
	Putter[K, V]
	Getter[K, V]
	Modifier[K, V]
	Deleter[K, V]

	Len(context.Context) (int, error)
	Cap(context.Context) (int, error)
}

// New returns a in-memory container, intended for prototyping and testing.
func New[K comparable, V any]() Container[K, V] {
	return make(mapWrap[K, V])
}
