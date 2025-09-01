/*
Package api contains all the configuration for the api that includes
the server and the routes
*/
package api

import "net/http"

type Myserver struct {
	*http.Server
}

func NewServer(config Config) Myserver {
	mux := NewRoutes()
	return Myserver{
		Server: &http.Server{
			Addr:    config.Addr,
			Handler: mux,
		},
	}
}

func (s Myserver) Start() error {
	err := s.ListenAndServe()
	return err
}
