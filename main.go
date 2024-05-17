package main

import (
	"fmt"
	"net/http"

	"github.com/mouhamedsylla/httprouter/httprouter"
)

func main() {
	router := httprouter.NewRouter()
	//router.Use(httprouter.CORSMiddleware)

	router.SetDirectory("/static/", "./static")
	router.Method(http.MethodGet).Handler("/static/", router.ServeStatic())
	router.Method(http.MethodGet).Handler("/api/message/private", privateStatic())
	router.Method(http.MethodGet).Handler("/api/message/private/:senderId", privateDynamic())
	router.Method(http.MethodGet).Handler("/api/message/private/:senderId/:receiverId", __privateDynamic())
	// router.Method(http.MethodGet, http.MethodPost).Middleware(httprouter.Mid1, httprouter.CORSMiddleware).Handler("/", Home())
	// router.Method(http.MethodGet).Middleware(httprouter.Mid2, httprouter.CORSMiddleware).Handler("/foo", FooPage())
	// router.Method(http.MethodGet).Middleware(httprouter.Mid3, httprouter.CORSMiddleware).Handler("/foo/bar", FooBarPage())
	// router.Method(http.MethodGet).Middleware(httprouter.CORSMiddleware).Handler("/test", Test())
	fmt.Println("tree:", router.Tree.Node.Child)

	for i, v := range router.Tree.Node.Child {
		fmt.Println(i, "  ---  ", v)
	}
	http.ListenAndServe(":8081", router)
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// server := http.Server{
	// 	Addr: ":8081",
	// }

	// go func() {
	// 	if err := server.ListenAndServe(); err != http.ErrServerClosed {
	// 		log.Fatalf("ListenAndServe(): %v", err)
	// 	}
	// }()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Arrêt du serveur en cours...")

	// ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Arrêt forcé du serveur : %v", err)
	// }
	// log.Println("Serveur arrêté")
}

func privateStatic() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get /api/message/private"))
	})
}

func privateDynamic() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get /api/message/private/:senderId"))
	})
}

func __privateDynamic() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/api/message/private/:senderId/:receiverId"))
	})
}

// func Home() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Get /"))
// 	})
// }

// func FooPage() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Get /foo"))
// 	})
// }

// func FooBarPage() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Get /foo/bar"))
// 	})
// }

// func Test() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		tmpl, err := template.ParseFiles("static/index.html")
// 		if err != nil {
// 			http.Error(w, "Ressource Not Found", http.StatusNotFound)
// 			return
// 		}

// 		tmpl.Execute(w, nil)
// 	})
// }
