package server

import (
	"fmt"
	"net/http"
)

func (handler *server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from GO Cast!")
}
