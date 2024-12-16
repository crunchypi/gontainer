package gontainer

import (
	"context"
	"errors"
)

var ErrSearchUpdater = errors.New("gontainer: failed search & update")
var ErrSearchDeleter = errors.New("gontainer: failed search & update")

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

// SearchUpdater represents something which searches and updates items.
type SearchUpdater[Q, U, R any] interface {
	SearchUpdate(ctx context.Context, filter Q, update U) (r R, err error)
}

// SearchDeleter represents something which searches and deletes items.
type SearchDeleter[Q, R any] interface {
	SearchDelete(ctx context.Context, filter Q) (r R, err error)
}
