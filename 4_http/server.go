package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About page, %s %s", r.URL.Query().Get("foo"), r.URL.Query().Get("bar"))
	})

	http.ListenAndServe(":8989", nil)
}
