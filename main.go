package main

import "net/http"

func main() {
	r := router.New()
	http.ListenAndServe(":8080", r.RootHandler())
}
