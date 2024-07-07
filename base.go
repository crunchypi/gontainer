package gontainer

import (
	"context"
	"errors"
)

var ErrPut = errors.New("gontainer: failed put")

// Putter represents something which stores a value.
type Putter[K comparable, V any] interface {
	Put(ctx context.Context, key K, val V) (err error)
}
