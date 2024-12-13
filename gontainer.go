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
var ErrMod = errors.New("gontainer: failed mod")
var ErrDel = errors.New("gontainer: failed del")

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

// -----------------------------------------------------------------------------
// Modifier.
// -----------------------------------------------------------------------------

// Modifier represents something which modifies a stored value.
type Modifier[K comparable, V any] interface {
	Mod(ctx context.Context, key K, rcv func(v V) V) (err error)
}

// ModifierImpl lets you implement Modifier with a function. The call to Mod is
// simply forwarded to the internal function "Impl".
//
// Example (interactive):
//   - https://go.dev/play/p/_zWjUTLFFwX
type ModifierImpl[K comparable, V any] struct {
	Impl func(
		ctx context.Context,
		key K,
		rcv func(v V) V,
	) (
		err error,
	)
}

// Mod implements Modifier by forwarding the call to the internal "Impl".
func (impl ModifierImpl[K, V]) Mod(
	ctx context.Context,
	key K,
	rcv func(v V) V,
) (
	err error,
) {
	if impl.Impl == nil {
		err = ErrImpl
		return
	}

	return impl.Impl(ctx, key, rcv)
}

// -----------------------------------------------------------------------------
// Deleter.
// -----------------------------------------------------------------------------

// Deleter represents something which deletes a stored value.
type Deleter[K comparable, V any] interface {
	Del(ctx context.Context, key K) (val V, err error)
}

// DeleterImpl lets you implement Deleter with a function. The call to Del is
// simply forwarded to the internal function "Impl".
//
// Example (interactive):
//   - https://go.dev/play/p/sEUi6zptniR
type DeleterImpl[K comparable, V any] struct {
	Impl func(
		ctx context.Context,
		key K,
	) (
		val V,
		err error,
	)
}

// Del implements Deleter by forwarding the call to the internal "Impl".
func (impl DeleterImpl[K, V]) Del(
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
