package main

import (
	"fmt"
	"reflect"
)

type Proxy[T any] struct {
	target T
}

func NewProxy[T any](target T) *Proxy[T] {
	return &Proxy[T]{target}
}

func (p *Proxy[T]) Invoke[V any](method func(T) V, args ...interface{}) V {
	in := make([]reflect.Value, len(args))
	for i := range args {
		in[i] = reflect.ValueOf(args[i])
	}

	out := reflect.ValueOf(method).Call(in)

	return out[0].Interface().(V)
}

type Hello interface {
	Greet(name string) string
}

type HelloImpl struct{}

func (h *HelloImpl) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	impl := &HelloImpl{}
	proxy := NewProxy(impl)

	hello := proxy.target.(Hello)
	fmt.Println(hello.Greet("world"))
}
