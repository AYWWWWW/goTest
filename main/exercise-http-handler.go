package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, s)
}

func (st *Struct) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, st.Greeting, st.Punct, st.Who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
