package gontainer

import (
	"context"
	"errors"
)

var ErrPut = errors.New("gontainer: failed put")
var ErrGet = errors.New("gontainer: failed get")

// Putter represents something which stores a value.
type Putter[K comparable, V any] interface {
	Put(ctx context.Context, key K, val V) (err error)
}

// Getter represents someting which gets a stored value.
type Getter[K comparable, V any] interface {
	Get(ctx context.Context, key K) (val V, err error)
}
