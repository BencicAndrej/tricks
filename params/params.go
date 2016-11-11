// Package params package unifies the way you access and set route parametes into a context.
//
// The package also demonstrates how to use context.Context.Value() properly,
// taken from the "context" package examples.
package params

import "context"

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

const (
	paramsKey key = iota
)

// NewContext returns a new Context with Request URL parameters map.
func NewContext(ctx context.Context, params map[string]string) context.Context {
	return context.WithValue(ctx, paramsKey, params)
}

// AddToContext is a helper function for adding parametes one by one.
func AddToContext(ctx context.Context, key, value string) context.Context {
	params := FromContext(ctx)
	params[key] = value

	return NewContext(ctx, params)
}

// FromContext returns the Request URL parameters map from Context.
func FromContext(ctx context.Context) map[string]string {
	return ctx.Value(paramsKey).(map[string]string)
}
