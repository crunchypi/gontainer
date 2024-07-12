package gontainer

import "context"

type mapWrap[K comparable, V any] map[K]V

// Put implements Putter.
func (m mapWrap[K, V]) Put(ctx context.Context, k K, v V) (err error) {
	m[k] = v
	return
}

// Get implements Getter.
func (m mapWrap[K, V]) Get(ctx context.Context, k K) (v V, err error) {
	v, ok := m[k]
	if !ok {
		err = ErrGet
	}

	return
}

func (m mapWrap[K, V]) Mod(ctx context.Context, k K, f func(V) V) (err error) {
	err = ErrImpl
	return
}

func (m mapWrap[K, V]) Del(ctx context.Context, k K) (v V, err error) {
	err = ErrImpl
	return
}

func (m mapWrap[K, V]) Len(context.Context) (n int, err error) {
	err = ErrImpl
	return
}

func (m mapWrap[K, V]) Cap(context.Context) (n int, err error) {
	err = ErrImpl
	return
}
