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

var ErrSearcher = errors.New("gontainer: failed search")
var ErrSearchUpdater = errors.New("gontainer: failed search & update")
var ErrSearchDeleter = errors.New("gontainer: failed search & update")

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

// -----------------------------------------------------------------------------
// Searcher
// -----------------------------------------------------------------------------

// Searcher represents something which searches for a value using a filter.
type Searcher[Q, R any] interface {
	Search(ctx context.Context, filter Q) (r R, err error)
}

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
// SearchUpdater
// -----------------------------------------------------------------------------

// SearchUpdater represents something which searches and updates items.
type SearchUpdater[Q, U, R any] interface {
	SearchUpdate(ctx context.Context, filter Q, update U) (r R, err error)
}

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
// SearchDeleter
// -----------------------------------------------------------------------------

// SearchDeleter represents something which searches and deletes items.
type SearchDeleter[Q, R any] interface {
	SearchDelete(ctx context.Context, filter Q) (r R, err error)
}

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

// -----------------------------------------------------------------------------
// Container.
// -----------------------------------------------------------------------------

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
