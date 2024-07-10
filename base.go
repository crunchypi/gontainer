package gontainer

import (
	"context"
	"errors"
)

var ErrPut = errors.New("gontainer: failed put")
var ErrGet = errors.New("gontainer: failed get")
var ErrMod = errors.New("gontainer: failed mod")
var ErrDel = errors.New("gontainer: failed del")

// Putter represents something which stores a value.
type Putter[K comparable, V any] interface {
	Put(ctx context.Context, key K, val V) (err error)
}

// Getter represents someting which gets a stored value.
type Getter[K comparable, V any] interface {
	Get(ctx context.Context, key K) (val V, err error)
}

// Modifier represents something which modifies a stored value.
type Modifier[K comparable, V any] interface {
	Mod(ctx context.Context, key K, rcv func(v V) V) (err error)
}

// Deleter represents something which deletes a stored value.
type Deleter[K comparable, V any] interface {
	Del(ctx context.Context, key K) (val V, err error)
}
