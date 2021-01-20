package service

import (
	"fmt"
	"net/http"
)

func Serve(mux *http.ServeMux, port string) error {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	return srv.ListenAndServe()
}
