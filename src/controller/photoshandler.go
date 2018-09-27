package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"

	"github.com/photosgallery/src/middleware"
	"github.com/photosgallery/src/model"
)

const (
	STATIC_DIR = "/tmp"
)

// InsertPhoto handles the new photos
func InsertPhoto() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(320 << 20)
		photo := model.Photo{}

		photo.ID = bson.NewObjectIdWithTime(time.Now())
		photo.Title = r.FormValue("title")
		photo.Likes = 0
		photo.Created = time.Now()

		conn, err := middleware.ExtractSession(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		file, handler, err := r.FormFile("filename")
		filename := fmt.Sprintf("%v/%v", "/tmp/", handler.Filename)
		photo.Filename = handler.Filename
		defer file.Close()
		dest, _ := os.Create(filename)
		defer dest.Close()
		io.Copy(dest, file)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		dao := model.PhotoDao{DB: conn}
		if err := dao.Insert(photo); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(photo)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photo)
	})
}

// GetPhotos retrieves all photos in the database
func GetPhotos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := middleware.ExtractSession(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		dao := model.PhotoDao{DB: conn}
		photos, err := dao.Read()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photos)
	})
}

/// FindPhoto
func FindPhoto() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := middleware.ExtractSession(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		params := mux.Vars(r)
		dao := model.PhotoDao{DB: conn}
		photo, err := dao.FindByID(params["id"])

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(photo)
	})
}

func DropPhoto() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := middleware.ExtractSession(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		param := mux.Vars(r)

		dao := model.PhotoDao{DB: conn}

		photo, err := dao.FindByID(param["id"])

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		err = os.Remove(STATIC_DIR + "/" + photo.Filename)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err := dao.Remove(param["id"]); err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}
