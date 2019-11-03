package main

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
	"go.uber.org/zap/zaptest"
)

func TestHandleTax(t *testing.T) {
	r := mux.NewRouter()
	taxDataDir, _ := filepath.Abs("./data/")
	logger := zaptest.NewLogger(t)
	defer logger.Sync()
	s := server{
		router:     r,
		taxDataDir: taxDataDir,
		logger:     logger,
	}
	s.registerRoutes()

	type statusCodeTest struct {
		path             string
		expectStatusCode int
	}

	statusCodeTests := []statusCodeTest{
		statusCodeTest{path: "/tax/notexist", expectStatusCode: http.StatusNotFound},
		statusCodeTest{path: "/tax/england", expectStatusCode: http.StatusOK},
	}

	for i, v := range statusCodeTests {
		r, _ := http.NewRequest("GET", v.path, nil)
		w := httptest.NewRecorder()
		s.ServeHTTP(w, r)
		res := w.Result()
		if res.StatusCode != v.expectStatusCode {
			t.Errorf("%d: %s %s failed: expected %d, got %d", i, "GET", v.path, v.expectStatusCode, res.StatusCode)
		}
	}
}
