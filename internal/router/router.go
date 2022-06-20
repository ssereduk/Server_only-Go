package router

import "net/http"

type Router struct{}

func New() *Router {
	return &Router{}
}

func (r *Router) RootHandler() http.Handler {
	return
}
