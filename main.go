package main

import (
	"net/http"
	"github.com/mouhamedsylla/httprouter"
)

func main() {
	router := httprouter.NewRouter()
	router.Use(httprouter.CORSMiddleware())

	router.Method(http.MethodGet).Handler("/", Home())
	router.Method(http.MethodPost).Handler("/foo", FooPage())
	router.Method(http.MethodGet).Handler("/foo/bar", FooBarPage())

	http.ListenAndServe(":8080", router)
}

func Home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get /"))
	})
}

func FooPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get /foo"))
	})
}

func FooBarPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get /foo/bar"))
	})
}
