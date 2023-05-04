package context

import (
	"context"
	"fmt"
)

type key string

func keyForValue(value any) key {
	return key(fmt.Sprintf("%T", value))
}

// Alternatively, we could use generics but it doesn't help us in this
// particular case.
//
// func WithValue[T any](ctx context.Context, value T) context.Context {
func WithValue(ctx context.Context, value any) context.Context {
	return context.WithValue(ctx, keyForValue(value), value)
}

func Value[T any](ctx context.Context) (T, error) {
	var tmp T
	val, ok := ctx.Value(keyForValue(tmp)).(T)
	if !ok {
		return tmp, fmt.Errorf("no value found")
	}
	return val, nil
}
