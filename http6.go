package main

import (
	"os"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handler("/", &myHandler{})
	mux.HandleFunc("/heelo", sayHello)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle(
		"/static/",
		http.StripPrefix(http.Dir(wd)))

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}

}


type myHandler struct {

}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString()
}