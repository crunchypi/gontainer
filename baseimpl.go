package gontainer

import (
	"context"
	"errors"
)

var ErrImpl = errors.New("gontainer: used interface without an implementation")

// -----------------------------------------------------------------------------
// Impl for Putter
// -----------------------------------------------------------------------------

// PutterImpl lets you implement Putter with a function. The call to Put is
// simply forwarded to the internal function "Impl".
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
// Impl for Getter.
// -----------------------------------------------------------------------------

// GetterImpl lets you implement Getter with a function. The call to Get is
// simply forwarded to the internal function "Impl".
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
// Impl for Modifier.
// -----------------------------------------------------------------------------

// ModifierImpl lets you implement Modifier with a function. The call to Mod is
// simply forwarded to the internal function "Impl".
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
// Impl for Deleter.
// -----------------------------------------------------------------------------

// DeleterImpl lets you implement Deleter with a function. The call to Del is
// simply forwarded to the internal function "Impl".
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
