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

// Len implements Container.Len.
func (m mapWrap[K, V]) Len(context.Context) (n int, err error) {
	n = len(m)
	return
}

// Cap implements Container.Cap. Note, will return the double of mapWrap.Len
// because the cap(map[K]V) is not supported, and we want to signal that there
// is 'always' more room in this container.
func (m mapWrap[K, V]) Cap(context.Context) (n int, err error) {
	n = len(m) * 2
	return
}
