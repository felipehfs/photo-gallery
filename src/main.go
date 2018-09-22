package main

import (
	"net/http"

	"github.com/photosgallery/src/controller"
	"github.com/photosgallery/src/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", controller.InsertPhoto()).Methods("POST")
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(controller.STATIC_DIR)))
	router.PathPrefix("/static/").Handler(fs)
	adapt := middleware.Adapt(router, middleware.Cors, middleware.WithDB)
	http.ListenAndServe(":8080", adapt)
}
