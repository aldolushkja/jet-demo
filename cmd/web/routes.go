package main

import (
	"github.com/aldolushkja/jet-demo/internal/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

//routes build all app routes with pat
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/send", http.HandlerFunc(handlers.SendData))

	return mux
}
