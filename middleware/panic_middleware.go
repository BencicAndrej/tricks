package middleware

import (
	"errors"
	"net/http"
)

// ErrUnknown is the error provided to PanicHandler if
// the recovered error does not match a regular error or a string.
var ErrUnknown = errors.New("Unkown error")

// PanicHandler executes in case of a panic during a request.
//
// Basic usecases for a PanicHandler are to notify the user that an error happened,
// log the panic error for later inverstigation, etc...
type PanicHandler func(w http.ResponseWriter, r *http.Request, err error)

// DefaultPanicHandler informs the user that the request has failed and
// does not act on the error that caused the panic.
func DefaultPanicHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// PanicMiddleware catches panics and provides the underlying error value
// to an optional callback.
func PanicMiddleware(ph PanicHandler) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err error
			defer func() {
				if rec := recover(); rec != nil {
					switch t := rec.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = ErrUnknown
					}

					ph(w, r, err)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}
