package tricks

import (
	"net/http"

	"github.com/bencicandrej/tricks/params"

	"github.com/julienschmidt/httprouter"
)

// HTTPRouterAdapter tranforms a regular http.Handler into httprouter.Handle.
func HTTPRouterAdapter(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := params.NewContext(r.Context(), paramsToMap(p))

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

func paramsToMap(params httprouter.Params) map[string]string {
	paramsMap := make(map[string]string)
	for _, param := range params {
		paramsMap[param.Key] = param.Value
	}

	return paramsMap
}
