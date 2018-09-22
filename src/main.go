package main

import (
	"net/http"

	"github.com/photosgallery/src/controller"
	"github.com/photosgallery/src/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/api/photos", controller.InsertPhoto()).Methods("POST")
	router.Handle("/api/photos", controller.GetPhotos()).Methods("GET")
	router.Handle("/api/photos/{id}", controller.FindPhoto()).Methods("GET")
	router.Handle("/api/photos/{id}", controller.DropPhoto()).Methods("DELETE")
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(controller.STATIC_DIR)))
	router.PathPrefix("/static/").Handler(fs)

	adapt := middleware.Adapt(router, middleware.Cors, middleware.WithDB)
	http.ListenAndServe(":8080", adapt)
}
