# gontainer
Generic container interfaces and decorators for Go

Index
- [Core interfaces](#core-interfaces)
- [Errors](#errors)
- [Impl pattern](#impl-pattern)



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



## Errors
```go
var ErrPut = errors.New("gontainer: failed put")
var ErrGet = errors.New("gontainer: failed get")
var ErrMod = errors.New("gontainer: failed mod")
var ErrDel = errors.New("gontainer: failed del")

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
