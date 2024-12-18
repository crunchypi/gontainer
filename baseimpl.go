package gontainer

import (
	"context"
)

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
