package main

import (
	"net/http"

	"github.com/ssereduk/Server_only-Go/tree/main/internal/router"
)

func main() {
	r := router.New()
	http.ListenAndServe(":8080", r.RootHandler())
}
