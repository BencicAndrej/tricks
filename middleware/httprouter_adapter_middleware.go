package middleware

import (
	"net/http"

	"github.com/bencicandrej/tricks"

	"github.com/julienschmidt/httprouter"
)

// DoHTTPRouter is a proxy to Stack.Do() method with an additional adapter
// for HTTPRouter package.
func (stack Stack) DoHTTPRouter(handler http.Handler) httprouter.Handle {
	return tricks.HTTPRouterAdapter(stack.Do(handler))
}
