package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
		var photo model.Photo

		photo.ID = bson.NewObjectIdWithTime(time.Now())
		photo.Title = r.FormValue("title")
		photo.Likes = 0
		photo.Created = time.Now()
		dest, err := uploadFile(r, "picture", STATIC_DIR)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		photo.URL = dest

		conn, err := middleware.ExtractSession(r)
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

func uploadFile(r *http.Request, field, dirname string) (string, error) {
	file, handler, err := r.FormFile(field)
	defer file.Close()
	if err != nil {
		return "", err
	}
	log.Println(handler.Filename)
	filename := fmt.Sprintf("%v/%v", dirname, handler.Filename)
	dest, err := os.Create(filename)
	io.Copy(dest, file)
	defer dest.Close()
	if err != nil {
		return "", err
	}
	return filename, nil
}
