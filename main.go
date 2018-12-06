package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"api-chi/api/todo"
)

func Routes() *chi.Mux{
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), //set content-type headers as application/json
		middleware.Logger, //log API request calls
		middleware.DefaultCompress, //compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, //redirect slashes to slash URL versions
		middleware.Recoverer, //recover from panics without crashing server  
	)

	router.Route("/v1", func(r chi.Router){
		r.Mount("/api/todo", todo.Routes())
	})

	// router.Route("/test", func(r chi.Router){
	// 	log.Printf("lock & lock")
	// })

	// log.Printf("LOG ME")
	// log.Printf("WERE")
	return router	
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func (http.Handler)http.Handler) error {
		log.Printf("%s %s\n", method, route) //walk and print out all routes
		return nil	
	}
	if err := chi.Walk(router, walkFunc); err != nil{
		log.Panicf("Logging err: %s\n") //panic if there is an error
	}

	log.Fatal(http.ListenAndServe(":8082", router)) //note, the port is usually gottem from
}

/* 
	$ go get github.com/go-chi
	or simply
	$ go get

	main source: https://itnext.io/structuring-a-production-grade-rest-api-in-golang-c0229b3feedc

	Middlewares are simply functions which get called before the handler is called, 
	so they are used for preprocessing requests, authentication, etc

*/