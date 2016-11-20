package middleware

import (
	"net/http"

	"github.com/bencicandrej/tricks"

	"github.com/julienschmidt/httprouter"
)

// Middleware is a function that takes a handler,
// is expected to do some work before or after provided
// handler and returns its handler.
type Middleware func(next http.Handler) http.Handler

// Stack is a list of Middleware.
type Stack struct {
	middleware []Middleware
}

// NewStack creates a MiddlewareStack with arbitrary
// number or middleware.
func NewStack(middleware ...Middleware) Stack {
	return Stack{
		middleware: append([]Middleware{}, middleware...),
	}
}

// Do adds the last Handler and returns the final,
// net/http compatible http.Handler.
func (stack Stack) Do(handler http.Handler) http.Handler {
	// Loop through all middleware backwards and construct
	// the final middleware.
	for i := range stack.middleware {
		handler = stack.middleware[len(stack.middleware)-1-i](handler)
	}

	return handler
}

// DoFunc is a shorthand of stack.Do(stack.HandlerFunc(f))
func (stack Stack) DoFunc(f http.HandlerFunc) http.Handler {
	return stack.Do(http.HandlerFunc(f))
}

// DoHTTPRouter is a proxy to Stack.Do() method with an additional adapter
// for HTTPRouter package.
func (stack Stack) DoHTTPRouter(handler http.Handler) httprouter.Handle {
	return tricks.HTTPRouterAdapter(stack.Do(handler))
}

// Append extends a stack, adding the specified middleware
// as the last ones in the request flow.
//
// Append returns a new stack, leaving the original one untouched.
func (stack Stack) Append(middleware ...Middleware) Stack {
	newMiddleware := make([]Middleware, len(stack.middleware)+len(middleware))
	copy(newMiddleware, stack.middleware)
	copy(newMiddleware[len(stack.middleware):], middleware)

	return NewStack(newMiddleware...)
}

// Extend extends a stack by adding the specified stack
// as the last one in the request flow.
//
// Extend returns a new stack, leaving the original one untouched.
func (stack Stack) Extend(newStack Stack) Stack {
	return stack.Append(newStack.middleware...)
}
