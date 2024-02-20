package main

import (
	"net/http"
	"text/template"

	"github.com/mouhamedsylla/httprouter/httprouter"
)

func main() {
	router := httprouter.NewRouter()
	//router.Use(httprouter.CORSMiddleware)

	router.SetDirectory("/static/", "./static")
	router.Method(http.MethodGet).Handler("/static/", router.ServeStatic())
	router.Method(http.MethodGet, http.MethodPost).Middleware(httprouter.Mid1, httprouter.CORSMiddleware).Handler("/", Home())
	router.Method(http.MethodGet).Middleware(httprouter.Mid2, httprouter.CORSMiddleware).Handler("/foo", FooPage())
	router.Method(http.MethodGet).Middleware(httprouter.Mid3, httprouter.CORSMiddleware).Handler("/foo/bar", FooBarPage())
	router.Method(http.MethodGet).Middleware(httprouter.CORSMiddleware).Handler("/test", Test())

	http.ListenAndServe(":8081", router)
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

func Test() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, "Ressource Not Found", http.StatusNotFound)
			return
		}

		tmpl.Execute(w, nil)
	})
}