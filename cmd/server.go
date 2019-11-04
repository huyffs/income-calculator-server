package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type responseType string

const (
	taxData responseType = "TAX_DATA"
)

type server struct {
	router     *mux.Router
	taxDataDir string
	logger     *zap.Logger
}

func (s *server) NotFound(w http.ResponseWriter, r *http.Request, e error) {
	s.logger.Sugar().Warn(e)
	http.NotFound(w, r)
}

func (s *server) InternalServerError(w http.ResponseWriter, e error) {
	s.logger.Sugar().Warn(e)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
