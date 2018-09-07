package main

import (
	"net/http"
	"log"
	"time"
	"io"
)


var mux map[string]func(w http.ResponseWriter, r *http.Request)


func main() {
	server := http.Server{
		Addr:              ":8080",
		Handler:           &myHandler{},
		ReadHeaderTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye


	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "url: "+r.URL.String())
}


func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "url: "+r.URL.String())
}
