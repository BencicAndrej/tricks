package middleware_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/bencicandrej/tricks/middleware"
)

func TestPanicMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("Hold up!")
	})

	stack := middleware.NewStack(middleware.PanicMiddleware())

	ts := httptest.NewServer(stack.Do(handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("request failed with error: %v", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("failed reading response body: %v", err)
		return
	}

	responseBody := strings.TrimRight(string(body), "\n\t ")

	if res.StatusCode != http.StatusInternalServerError || responseBody != http.StatusText(http.StatusInternalServerError) {
		t.Errorf(
			"Got resposne [%d] %q, wanted [%d] %q",
			res.StatusCode,
			responseBody,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}
}
