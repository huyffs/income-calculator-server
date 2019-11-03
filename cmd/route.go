package main

func (s *server) registerRoutes() {
	s.router.HandleFunc("/tax/{region}", s.logMi(s.handleTax()))
}
