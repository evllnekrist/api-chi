package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/tonyalaribe/todoapi/basestructure/features/todo"
)

func Routes() *chi.Mux{
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), //set content-type headers as application/json
		middleware.Logger, //log API request calls
		middleware.DefaultCompress, //compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // redirect slashes to slash URL versions
		middleware.Recoverer, //  
	)
}


/* 
	$ go get github.com/go-chi
	or simply
	$ go get

	main source: https://itnext.io/structuring-a-production-grade-rest-api-in-golang-c0229b3feedc
*/