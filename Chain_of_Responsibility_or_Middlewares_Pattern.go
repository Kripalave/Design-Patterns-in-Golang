package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

/*
As you can see, a middleware is a function that:
accepts a http.HandlerFunc as an argument;
returns a http.HandlerFunc;
calls the http.handlerFunc passed in.
With this basic technique you can "chain" as many middlewares as you like:
*/

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/foo", middleware(handler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/MODIFIED"

		// Run the next handler
		next.ServeHTTP(w, r)
	}
}
