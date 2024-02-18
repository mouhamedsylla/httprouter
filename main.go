package main

import (
	"net/http"

	"github.com/mouhamedsylla/httprouter/httprouter"
)

func main() {
	router := httprouter.NewRouter()
	//router.Use(httprouter.CORSMiddleware)

	router.Method(http.MethodGet).Middleware(httprouter.CORSMiddleware, httprouter.Mid1).Handler("/", Home())
	router.Method(http.MethodGet).Middleware(httprouter.CORSMiddleware, httprouter.Mid2).Handler("/foo", FooPage())
	router.Method(http.MethodGet).Middleware(httprouter.CORSMiddleware, httprouter.Mid3).Handler("/foo/bar", FooBarPage())

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
