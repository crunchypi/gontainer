package gontainer

import (
	"context"
)

// -----------------------------------------------------------------------------
// Impl for Getter.
// -----------------------------------------------------------------------------

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
// Impl for Modifier.
// -----------------------------------------------------------------------------

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
// Impl for Deleter.
// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------
// Impl for Container.
// -----------------------------------------------------------------------------

// ContainerImpl lets you implement Container with functions. It groups
// PutterImpl, GetterImpl, ModifierImpl and DeleterImpl, so docs for those may
// be useful to read. Also, similarly you may implement the Len and Cap func.
//
// Example (interactive):
//   - https://go.dev/play/p/QdFBbTL5v3_E
type ContainerImpl[K comparable, V any] struct {
	PutterImpl[K, V]
	GetterImpl[K, V]
	ModifierImpl[K, V]
	DeleterImpl[K, V]

	ImplLen func(ctx context.Context) (int, error)
	ImplCap func(ctx context.Context) (int, error)
}

// Len implements Container.Len by forwarding the call to the internal "ImplLen".
func (impl ContainerImpl[K, V]) Len(
	ctx context.Context,
) (
	n int,
	err error,
) {
	if impl.ImplLen == nil {
		err = ErrImpl
		return
	}

	return impl.ImplLen(ctx)
}

// Cap implements Container.Cap by forwarding the call to the internal "ImplCap".
func (impl ContainerImpl[K, V]) Cap(
	ctx context.Context,
) (
	n int,
	err error,
) {
	if impl.ImplCap == nil {
		err = ErrImpl
		return
	}

	return impl.ImplCap(ctx)
}

// -----------------------------------------------------------------------------
// Impl for Searcher
// -----------------------------------------------------------------------------

// SearcherImpl lets you implement Searcher with a function. The call to Search
// is simply forwarded to the internal function "Impl".
//
// Example (interactive):
//   - https://go.dev/play/p/KuzLaYVfYct
type SearcherImpl[Q, R any] struct {
	Impl func(ctx context.Context, filter Q) (r R, err error)
}

// Search implements Searcher.Search by forwarding the call to the internal "Impl".
func (impl SearcherImpl[Q, R]) Search(
	ctx context.Context,
	filter Q,
) (
	r R,
	err error,
) {
	if impl.Impl == nil {
		err = ErrImpl
		return
	}

	return impl.Impl(ctx, filter)
}

// -----------------------------------------------------------------------------
// Impl for SearchUpdater
// -----------------------------------------------------------------------------

// SearchUpdaterImpl lets you implement SearchUpdater with a function. The call
// to SearchUpdate is simply forwarded to the internal function "Impl".
//
// Example (interactive):
//   - https://go.dev/play/p/-9AdaI2w4GJ
type SearchUpdaterImpl[Q, U, R any] struct {
	Impl func(ctx context.Context, filter Q, update U) (r R, err error)
}

// SearchUpdate implements SearchUpdater by forwarding to the internal "Impl".
func (impl SearchUpdaterImpl[Q, U, R]) SearchUpdate(
	ctx context.Context,
	filter Q,
	update U,
) (
	r R,
	err error,
) {
	if impl.Impl == nil {
		err = ErrImpl
		return
	}

	return impl.Impl(ctx, filter, update)
}

// -----------------------------------------------------------------------------
// Impl for SearchDeleter
// -----------------------------------------------------------------------------

// SearchDeleterImpl lets you implement SearchDeleter with a function. The call
// to SearchDelete is simply forwarded to the internal function "Impl".
//
// Example (interactive):
//   - https://go.dev/play/p/sJC4P3nR_ML
type SearchDeleterImpl[Q, R any] struct {
	Impl func(ctx context.Context, filter Q) (q R, err error)
}

// SearchDelete implements SearchDeleter by forwarding to the internal "Impl".
func (impl SearchDeleterImpl[Q, R]) SearchDelete(
	ctx context.Context,
	filter Q,
) (
	r R,
	err error,
) {
	if impl.Impl == nil {
		err = ErrImpl
		return
	}

	return impl.Impl(ctx, filter)
}
