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

// Mod implements Modifier. Note, will still do a write if "k" is not found.
func (m mapWrap[K, V]) Mod(ctx context.Context, k K, f func(V) V) (err error) {
	if f == nil {
		return
	}

	v, ok := m[k]
	if !ok {
		err = ErrMod
	}

	m[k] = f(v)
	return
}

// Del implements Deleter.
func (m mapWrap[K, V]) Del(ctx context.Context, k K) (v V, err error) {
	v, ok := m[k]
	if !ok {
		err = ErrDel
		return
	}

	delete(m, k)
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
