package main

import (
	"net/http"
	"log"
	"io"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "url: "+r.URL.String())

}
