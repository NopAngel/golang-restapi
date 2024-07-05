package server

import "net/http"

type Contry struct {
	Name     string
	Language string
}

var contries []*Contry = []*Contry{}

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
