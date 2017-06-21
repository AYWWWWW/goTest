package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Fprintf(w, "key:", k)
		fmt.Fprintf(w, "value:", v)
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func validateForm(r *http.Request) bool{
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		return false
	}
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
		return false
	}
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		return false
	}
	if m, _ := regexp.MatchString("^([w\.\_]{2,10}@(\w{1,}).([a-z]{2,4})$", r.Form.Get("email")); !m {
		return false
	}

}

type MyMux struct {
}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	switch r.URL.Path {
	case "/":
		sayhelloName(w, r)
		return
	case "/login":
		login(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
