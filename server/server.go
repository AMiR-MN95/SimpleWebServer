package server

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	counter int
}

func New() *server {
	return &server{
		counter: 0,
	}
}

func (s *server) Server(port int) {
	http.HandleFunc("/hello", s.hello)
	http.HandleFunc("/bmi", s.bmi)
	http.HandleFunc("/counter", s.incrementCounter)

	http.Handle("/sample", &sample{})

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type sample struct {
}

func (s *sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample handler")
}
