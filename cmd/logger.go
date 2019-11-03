package main

import (
	"net/http"
)

func (s *server) logMi(h http.HandlerFunc) http.HandlerFunc {
	sugar := s.logger.Sugar()
	return func(w http.ResponseWriter, r *http.Request) {
		sugar.Infow("begin", "method", r.Method, "path", r.URL.Path)
		defer sugar.Infow("end", "method", r.Method, "path", r.URL.Path)
		h(w, r)
	}
}
