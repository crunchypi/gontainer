# gontainer
Generic container interfaces and decorators for Go

Index
- [Core interfaces](#core-interfaces)
- [Errors](#errors)
- [Impl pattern](#impl-pattern)
- [Default](#default)



## Core interfaces

#### Putter
```go
// Putter represents something which stores a value.
type Putter[K comparable, V any] interface {
	Put(ctx context.Context, key K, val V) (err error)
}
```

#### Getter
```go
// Getter represents someting which gets a stored value.
type Getter[K comparable, V any] interface {
	Get(ctx context.Context, key K) (val V, err error)
}
```

#### Modifier
```go
// Modifier represents something which modifies a stored value.
type Modifier[K comparable, V any] interface {
	Mod(ctx context.Context, key K, rcv func(v V) V) (err error)
}
```

#### Deleter
```go
// Deleter represents something which deletes a stored value.
type Deleter[K comparable, V any] interface {
	Del(ctx context.Context, key K) (val V, err error)
}
```

#### Container
```go
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
```

#### Searcher
```go
// Searcher represents something which searches for a value using a filter.
type Searcher[Q, R any] interface {
	Search(ctx context.Context, filter Q) (r R, err error)
}
```

#### SearchUpdater
```go
// SearchUpdater represents something which searches and updates items.
type SearchUpdater[Q, U, R any] interface {
	SearchUpdate(ctx context.Context, filter Q, update U) (r R, err error)
}
```

#### SearchDeleter

```go
// SearchDeleter represents something which searches and deletes items.
type SearchDeleter[Q, R any] interface {
	SearchDelete(ctx context.Context, filter Q) (r R, err error)
}
```


## Errors
```go
var ErrPut = errors.New("gontainer: failed put")
var ErrGet = errors.New("gontainer: failed get")
var ErrMod = errors.New("gontainer: failed mod")
var ErrDel = errors.New("gontainer: failed del")

var ErrSearchFinder = errors.New("gontainer: failed search")
var ErrSearchUpdater = errors.New("gontainer: failed search & update")
var ErrSearchDeleter = errors.New("gontainer: failed search & update")

// See the next section.
var ErrImpl = errors.New("gontainer: used interface without an implementation")
```



## Impl pattern

The impl pattern allows you to implement an interface in a functional way, avoiding the tedium of defining structs which implement small interfaces. You simply define the function and place it inside an impl struct.

#### Impl for Putter.
```go
// PutterImpl lets you implement Putter with a function. The call to Put is
// simply forwarded to the internal function "Impl".
type PutterImpl[K comparable, V any] struct {
	Impl func(ctx context.Context, key K, val V) (err error)
}

// Put implements Putter by forwarding the call to the internal "Impl".
func (impl PutterImpl[K, V]) Put(ctx context.Context, key K, val V) (err error)
```

#### Impl for Getter.
```go
// GetterImpl lets you implement Getter with a function. The call to Get is
// simply forwarded to the internal function "Impl".
type GetterImpl[K comparable, V any] struct {
	Impl func(ctx context.Context, key K) (val V, err error,)
}

// Get implements Getter by forwarding the call to the internal "Impl".
func (impl GetterImpl[K, V]) Get(ctx context.Context, key K) (val V, err error)
```

#### Impl for Modifier.
```go
// ModifierImpl lets you implement Modifier with a function. The call to Mod is
// simply forwarded to the internal function "Impl".
type ModifierImpl[K comparable, V any] struct {
	Impl func(ctx context.Context, key K, rcv func(v V) V) (err error)
}

// Mod implements Modifier by forwarding the call to the internal "Impl".
func (impl ModifierImpl[K, V]) Mod(ctx context.Context, key K, rcv func(v V) V) (err error)
```

#### Impl for Deleter.
```go
// DeleterImpl lets you implement Deleter with a function. The call to Del is
// simply forwarded to the internal function "Impl".
type DeleterImpl[K comparable, V any] struct {
	Impl func(ctx context.Context, key K) (val V, err error)
}

// Del implements Deleter by forwarding the call to the internal "Impl".
func (impl DeleterImpl[K, V]) Del(ctx context.Context, key K) (val V, err error)
```

#### Impl for Container.
```go
// ContainerImpl lets you implement Container with functions. It groups
// PutterImpl, GetterImpl, ModifierImpl and DeleterImpl, so docs for those may
// be useful to read. Also, similarly you may implement the Len and Cap func.
type ContainerImpl[K comparable, V any] struct {
	PutterImpl[K, V]
	GetterImpl[K, V]
	ModifierImpl[K, V]
	DeleterImpl[K, V]

	ImplLen func(ctx context.Context) (int, error)
	ImplCap func(ctx context.Context) (int, error)
}

// Len implements Container.Len by forwarding the call to the internal "ImplLen".
func (impl ContainerImpl[K, V]) Len(ctx context.Context) (n int, err error)

// Cap implements Container.Cap by forwarding the call to the internal "ImplCap".
func (impl ContainerImpl[K, V]) Cap(ctx context.Context) (n int, err error)
```

#### Impl for Searcher
```go
// SearcherImpl lets you implement Searcher with a function. The call to Search
// is simply forwarded to the internal function "Impl".
type SearcherImpl[Q, R any] struct {
	Impl func(ctx context.Context, filter Q) (r R, err error)
}

// Search implements Searcher.Search by forwarding the call to the internal "Impl".
func (impl SearcherImpl[Q, R]) Search(ctx context.Context, filter Q) R, error) 
```

#### Impl for SearchUpdater
```go
// SearchUpdaterImpl lets you implement SearchUpdater with a function. The call
// to SearchUpdate is simply forwarded to the internal function "Impl".
type SearchUpdaterImpl[Q, U, R any] struct {
	Impl func(ctx context.Context, filter Q, update U) (r R, err error)
}

// SearchUpdate implements SearchUpdater by forwarding to the internal "Impl".
func (impl SearchUpdaterImpl[Q, U, R]) SearchUpdate(ctx context.Context, filter Q, update U,) (r R, err error) 
```

#### Impl for SearchDeleter
```go
// SearchDeleterImpl lets you implement SearchDeleter with a function. The call
// to SearchDelete is simply forwarded to the internal function "Impl".
type SearchDeleterImpl[Q, R any] struct {
	Impl func(ctx context.Context, filter Q) (q R, err error)
}

// SearchDelete implements SearchDeleter by forwarding to the internal "Impl".
func (impl SearchDeleterImpl[Q, R]) SearchDelete(ctx context.Context, filter Q) (r R, err error)
```


## Default
A default container is implemented and has the signature noted below. It is mainly intended for prototyping and testing, and is implemented as a `map[K]V`wrapped by the `Container` interface. The underlying implementation may be swapped in the future but the signature and behavior will most likely not.

Some notes:
- As `cap(map[K]V)` is not supported by the language, a call to `Cap` returns `Len` * 2
- `Mod`will run the callback and save the result even if the key does not exist.

```go
func New[K comparable, V any]() Container[K, V]
```