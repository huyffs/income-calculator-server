package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	r := mux.NewRouter()
	taxDataDir, err := filepath.Abs("./data/")
	if err != nil {
		return err
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer logger.Sync()
	s := server{
		router:     r,
		taxDataDir: taxDataDir,
		logger:     logger,
	}
	s.registerRoutes()
	return http.ListenAndServe(":8000", r)
}
