package gontainer

import (
	"context"
	"errors"
)

// -----------------------------------------------------------------------------
// Errors
// -----------------------------------------------------------------------------

var ErrPut = errors.New("gontainer: failed put")
var ErrGet = errors.New("gontainer: failed get")

var ErrImpl = errors.New("gontainer: used interface without an implementation")

// -----------------------------------------------------------------------------
// Putter
// -----------------------------------------------------------------------------

// Putter represents something which stores a value.
type Putter[K comparable, V any] interface {
	Put(ctx context.Context, key K, val V) (err error)
}

// PutterImpl lets you implement Putter with a function. The call to Put is
// simply forwarded to the internal function "Impl".
//
// Example (interactive):
//   - https://go.dev/play/p/3QuCteK2sU6
type PutterImpl[K comparable, V any] struct {
	Impl func(
		ctx context.Context,
		key K,
		val V,
	) (
		err error,
	)
}

// Put implements Putter by forwarding the call to the internal "Impl".
func (impl PutterImpl[K, V]) Put(
	ctx context.Context,
	key K,
	val V,
) (
	err error,
) {
	if impl.Impl == nil {
		err = ErrImpl
		return
	}

	return impl.Impl(ctx, key, val)
}

// -----------------------------------------------------------------------------
// Getter.
// -----------------------------------------------------------------------------

// Getter represents someting which gets a stored value.
type Getter[K comparable, V any] interface {
	Get(ctx context.Context, key K) (val V, err error)
}

// GetterImpl lets you implement Getter with a function. The call to Get is
// simply forwarded to the internal function "Impl".
//
// Example (interactive):
//   - https://go.dev/play/p/iNY6Lcf0Bmo
type GetterImpl[K comparable, V any] struct {
	Impl func(
		ctx context.Context,
		key K,
	) (
		val V,
		err error,
	)
}

// Get implements Getter by forwarding the call to the internal "Impl".
func (impl GetterImpl[K, V]) Get(
	ctx context.Context,
	key K,
) (
	val V,
	err error,
) {
	if impl.Impl == nil {
		err = ErrImpl
		return
	}

	return impl.Impl(ctx, key)
}
