package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// import "./utils"
import "hello"
var mu sync.Mutex
var count int

func main() {
	fmt.Fprintf("%s", hello.Proverb())
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	handlerGif := func(w http.ResponseWriter, r *http.Request) {
		// utils.Lissajous(w)
	}
	http.HandleFunc("/gif", handlerGif)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock();
}